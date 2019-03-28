package productplan

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"reflect"

	"github.com/google/go-querystring/query"
)

const (
	// Version used in the user-agent identification
	Version = "1.0"

	// userAgent represents the default user agent used
	// when no other user agent is set.
	defaultUserAgent = "productplan-go/" + Version

	apiVersion = "1"
)

// Client represents a client to the Productplan API.
type Client struct {
	// HTTPClient is the underlying HTTP client
	// used to communicate with the API.
	HTTPClient *http.Client

	// Credentials used for accessing the Productplan API
	Credentials Credentials

	// BaseURL for API requests.
	BaseURL string

	// UserAgent used when communicating with the Productplan API.
	UserAgent string

	Status   *StatusService
	Ideas    *IdeasService
	Roadmaps *RoadmapsService
	Bars     *BarsService

	// Set to true to output debugging logs during API calls
	Debug bool
}

// NewClient returns a new ProductPlan API client using the given credentials.
func NewClient(endpoint string, credentials Credentials) *Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	c := &Client{Credentials: credentials, HTTPClient: &http.Client{Transport: tr}, BaseURL: endpoint}
	c.Status = &StatusService{client: c}
	c.Ideas = &IdeasService{client: c}
	c.Roadmaps = &RoadmapsService{client: c}
	c.Bars = &BarsService{client: c}
	c.Debug = false
	return c
}

// NewRequest creates an API request.
// The path is expected to be a relative path and will be resolved
// according to the BaseURL of the Client. Paths should always be specified without a preceding slash.
func (c *Client) NewRequest(method, path string, payload interface{}) (*http.Request, error) {
	url := c.BaseURL + path

	body := new(bytes.Buffer)
	if payload != nil {
		err := json.NewEncoder(body).Encode(payload)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Api-Version", apiVersion)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", formatUserAgent(c.UserAgent))
	for key, value := range c.Credentials.Headers() {
		req.Header.Add(key, value)
	}

	return req, nil
}

func (c *Client) get(path string, obj interface{}) (*http.Response, error) {
	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	return c.Do(req, obj)
}

func (c *Client) post(path string, payload, obj interface{}) (*http.Response, error) {
	req, err := c.NewRequest("POST", path, payload)
	if err != nil {
		return nil, err
	}

	return c.Do(req, obj)
}

func (c *Client) patch(path string, payload, obj interface{}) (*http.Response, error) {
	req, err := c.NewRequest("PATCH", path, payload)
	if err != nil {
		return nil, err
	}

	return c.Do(req, obj)
}

// Do sends an API request and returns the API response.
//
// The API response is JSON decoded and stored in the value pointed by obj,
// or returned as an error if an API error has occurred.
// If obj implements the io.Writer interface, the raw response body will be written to obj,
// without attempting to decode it.
func (c *Client) Do(req *http.Request, obj interface{}) (*http.Response, error) {
	if c.Debug {
		log.Printf("Executing request (%v): %#v", req.URL, req)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if c.Debug {
		log.Printf("Response received: %#v", resp)
	}

	err = CheckResponse(resp)
	if err != nil {
		return resp, err
	}

	// If obj implements the io.Writer,
	// the response body is decoded into w.
	if obj != nil {
		if w, ok := obj.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(obj)
		}
	}

	return resp, err
}

// A Response represents an API response.
type Response struct {
	HTTPResponse *http.Response

	// If the response is paginated, the Pagination will store them.
	Pagination *Pagination `json:"pagination"`
}

// ListOptions contains the common options you can pass to a List method
// in order to control parameters such as paginations and page number.
type ListOptions struct {
	// Limit query given by string.
	Filters string `url:"filters"`

	// The page to return
	Page int `url:"page,omitempty"`

	// The number of entries to return per page
	Items int `url:"items,omitempty"`

	// The order criteria to sort the results.
	// The value is a comma-separated list of field[:direction],
	// eg. name | name:desc | name:desc,expiration:desc
	Order string `url:"order,omitempty"`
}

// Pagination If the response is paginated, Pagination represents the pagination information.
type Pagination struct {
	CurrentPage  int `json:"current_page"`
	PerPage      int `json:"per_page"`
	TotalPages   int `json:"total_pages"`
	TotalEntries int `json:"total_entries"`
}

// An ErrorResponse represents an API response that generated an error.
type ErrorResponse struct {
	Response
	Message string `json:"message"`
}

// Error implements the error interface.
func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %v %v",
		r.HTTPResponse.Request.Method, r.HTTPResponse.Request.URL,
		r.HTTPResponse.StatusCode, r.Message)
}

// CheckResponse checks the API response for errors, and returns them if present.
// A response is considered an error if the status code is different than 2xx. Specific requests
// may have additional requirements, but this is sufficient in most of the cases.
func CheckResponse(resp *http.Response) error {
	if code := resp.StatusCode; 200 <= code && code <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{}
	errorResponse.HTTPResponse = resp

	err := json.NewDecoder(resp.Body).Decode(errorResponse)
	if err != nil {
		return err
	}

	return errorResponse
}

// formatUserAgent builds the final user agent to use for HTTP requests.
func formatUserAgent(customUserAgent string) string {
	if customUserAgent == "" {
		return defaultUserAgent
	}

	return fmt.Sprintf("%s %s", defaultUserAgent, customUserAgent)
}

// addOptions adds the parameters in opt as URL query parameters to s.  opt
// must be a struct whose fields may contain "url" tags.
func addURLQueryOptions(path string, options interface{}) (string, error) {
	opt := reflect.ValueOf(options)

	// options is a pointer
	// return if the value of the pointer is nil,
	if opt.Kind() == reflect.Ptr && opt.IsNil() {
		return path, nil
	}

	// append the options to the URL
	u, err := url.Parse(path)
	if err != nil {
		return path, err
	}

	qs, err := query.Values(options)
	if err != nil {
		return path, err
	}

	uqs := u.Query()
	for k := range qs {
		uqs.Set(k, qs.Get(k))
	}
	u.RawQuery = uqs.Encode()
	return u.String(), nil
}
