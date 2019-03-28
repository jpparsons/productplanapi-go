package productplan

// StatusService handles communication with the status
// methods of the Productplan API.
type StatusService struct {
	client *Client
}

// Status represents the status of the API components
type Status struct {
	Application string `json:"application,omitempty"`
	Database    string `json:"database,omitempty"`
}

// Metadata represents application metadata in Productplan.
type Metadata struct {
	Application string `json:"application,omitempty"`
	Version     int    `json:"version,omitempty"`
}

// StatusResponse represents a response from an API method that returns an APIStatus struct.
type StatusResponse struct {
	Response
	Metadata
	Status `json:"status"`
}

// GetStatus get API status
func (s *StatusService) GetStatus() (*StatusResponse, error) {
	path := "/api/status"
	statusResponse := &StatusResponse{}

	resp, err := s.client.get(path, statusResponse)
	if err != nil {
		return nil, err
	}

	statusResponse.HTTPResponse = resp
	return statusResponse, nil
}
