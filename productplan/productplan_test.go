package productplan

import (
	"bufio"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	mux    *http.ServeMux
	client *Client
	server *httptest.Server
)

func setupMockServer() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	client = NewClient("localhost", NewOauthTokenCredentials("productplan-token"))
	client.BaseURL = server.URL
}

func teardownMockServer() {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; want != got {
		t.Errorf("Request METHOD expected to be `%v`, got `%v`", want, got)
	}
}

func testHeader(t *testing.T, r *http.Request, name, want string) {
	if got := r.Header.Get(name); want != got {
		t.Errorf("Request() %v expected to be `%#v`, got `%#v`", name, want, got)
	}
}

func testHeaders(t *testing.T, r *http.Request) {
	testHeader(t, r, "Accept", "application/json")
	testHeader(t, r, "User-Agent", defaultUserAgent)
}

func readHTTPFixture(t *testing.T, filename string) string {
	data, err := ioutil.ReadFile("../fixtures.http" + filename)
	if err != nil {
		t.Fatalf("Unable to read HTTP fixture: %v", err)
	}

	// Terrible hack
	// Some fixtures have \n and not \r\n

	// Terrible hack
	s := string(data[:])
	s = strings.Replace(s, "Transfer-Encoding: chunked\n", "", -1)
	s = strings.Replace(s, "Transfer-Encoding: chunked\r\n", "", -1)

	return s
}

func httpResponseFixture(t *testing.T, filename string) *http.Response {
	resp, err := http.ReadResponse(bufio.NewReader(strings.NewReader(readHTTPFixture(t, filename))), nil)
	if err != nil {
		t.Fatalf("Unable to create http.Response from fixture: %v", err)
	}

	return resp
}

func TestNewClient(t *testing.T) {
	c := NewClient("localhost", NewOauthTokenCredentials("productplan-token"))

	if c.BaseURL != "localhost" {
		t.Errorf("NewClient BaseURL = %v, want %v", c.BaseURL, "localhost")
	}
}

func TestClient_NewRequest(t *testing.T) {
	c := NewClient("localhost", NewOauthTokenCredentials("productplan-token"))
	c.BaseURL = "https://go.example.com"

	inURL, outURL := "/foo", "https://go.example.com/foo"
	req, _ := c.NewRequest("GET", inURL, nil)

	// test that relative URL was expanded with the proper BaseURL
	if req.URL.String() != outURL {
		t.Errorf("NewRequest(%v) URL = %v, want %v", inURL, req.URL, outURL)
	}

	// test that default user-agent is attached to the request
	ua := req.Header.Get("User-Agent")
	if ua != defaultUserAgent {
		t.Errorf("NewRequest() User-Agent = %v, want %v", ua, defaultUserAgent)
	}

	v := req.Header.Get("X-Api-Version")
	if v != apiVersion {
		t.Errorf("NewRequest() X-Api-Version = %v, want %v", v, apiVersion)
	}
}
