package productplan

import (
	"fmt"
)

// IdeasService handles communication with the idea
// methods of the Productplan API.
type IdeasService struct {
	client *Client
}

// Ideas represents an idea in ProductPlan
type Ideas struct {
	Href           string            `json:"href,omitempty"`
	ID             int               `json:"id,omitempty"`
	Name           string            `json:"name"`
	Description    string            `json:"description,omitempty"`
	StrategicValue string            `json:"strategic_value,omitempty"`
	Notes          string            `json:"notes,omitempty"`
	PercentDone    int               `json:"percent_done,omitempty"`
	Effort         int               `json:"effort,omitempty"`
	Tags           []string          `json:"tags,omitempty"`
	Fields         map[string]string `json:"fields,omitempty"`
	Timestamps     *Timestamps       `json:"timestamps,omitempty"`
	IdeaLinks      *IdeaLinks        `json:"links,omitempty"`
}

// IdeaLinks on an idea
type IdeaLinks struct {
	Roadmap       map[string]string `json:"roadmap"`
	ExternalLinks map[string]string `json:"external_links"`
}

// IdeasResponse represents a response from an API method that returns an Ideas struct.
type IdeasResponse struct {
	Response
	Ideas
}

// Show an idea
func (s *IdeasService) Show(id string) (*IdeasResponse, error) {
	path := fmt.Sprintf("/api/ideas/%v", id)
	ideasResponse := &IdeasResponse{}

	resp, err := s.client.get(path, ideasResponse)
	if err != nil {
		return nil, err
	}

	ideasResponse.HTTPResponse = resp
	return ideasResponse, nil
}

// TODO: index show all ideas paginated
