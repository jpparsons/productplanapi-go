package productplan

import (
	"fmt"
	"io"
)

// BarsService handles communication with the bar
// methods of the Productplan API.
type BarsService struct {
	client *Client
}

// Bar represents a bar
type Bar struct {
	Href           string            `json:"href"`
	ID             int               `json:"id"`
	Name           string            `json:"name"`
	StartDate      string            `json:"start_date,omitempty"`
	EndDate        string            `json:"end_date,omitempty"`
	Description    string            `json:"description,omitempty"`
	StrategicValue string            `json:"strategic_value,omitempty"`
	Notes          string            `json:"notes,omitempty"`
	PercentDone    int               `json:"percent_done,omitempty"`
	Effort         int               `json:"effort,omitempty"`
	Tags           []string          `json:"tags,omitempty"`
	Fields         map[string]string `json:"fields,omitempty"`
	Timestamps     `json:"timestamps"`
	BarLinks       `json:"links,omitempty"`
}

// BarLinks on a bar
type BarLinks struct {
	Roadmap       map[string]string `json:"roadmap,omitempty"`
	ParentBar     map[string]string `json:"parent_bar,omitempty"`
	ChildBars     map[string]string `json:"child_bars,omitempty"`
	ExternalLinks map[string]string `json:"external_links,omitempty"`
}

// UpdateBar bar attributes to perform an update
type UpdateBar struct {
	Name           string            `json:"name,omitempty"`
	StartDate      string            `json:"start_date,omitempty"`
	EndDate        string            `json:"end_date,omitempty"`
	Description    string            `json:"description,omitempty"`
	StrategicValue string            `json:"strategic_value,omitempty"`
	Notes          string            `json:"notes,omitempty"`
	PercentDone    int               `json:"percent_done,omitempty"`
	Effort         int               `json:"effort,omitempty"`
	Tags           []string          `json:"tags,omitempty"`
	Fields         map[string]string `json:"fields,omitempty"`
}

// BarsResponse represents a response from an API method that returns a bars struct.
type BarsResponse struct {
	Response
	Bar
}

// ListBars get a list of bars owned by the authenticated user
func (s *BarsService) ListBars(options *ListOptions) (*[]BarsResponse, error) {
	path := "/api/bars"
	var barsResponse *[]BarsResponse

	path, err := addURLQueryOptions(path, options)
	if err != nil {
		return nil, err
	}

	_, err = s.client.get(path, &barsResponse)
	if err != nil {
		return nil, err
	}

	return barsResponse, nil
}

// UpdateBar updates a bar
func (s *BarsService) UpdateBar(id int, barAttributes interface{}) (*BarsResponse, error) {
	path := fmt.Sprintf("/api/bars/%v", id)
	barsResponse := &BarsResponse{}

	resp, err := s.client.patch(path, barAttributes, barsResponse)

	// update does not return any response data, so expect the response decode to be EOF
	if err != nil && err != io.EOF {
		return nil, err
	}

	barsResponse.HTTPResponse = resp
	return barsResponse, nil
}
