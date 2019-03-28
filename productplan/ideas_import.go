package productplan

import (
	"io"
)

// IdeaImportRoadmap represents a roadmap for the payload
type IdeaImportRoadmap struct {
	ID int `json:"id"`
}

// IdeasImportAttributes represent .
type IdeasImportAttributes struct {
	IdeaImportRoadmap `json:"roadmap"`
	Ideas             []Ideas `json:"ideas"`
}

// IdeasImportResponse represents a response from an Ideas import.
type IdeasImportResponse struct {
	Response
}

// Import handles ideas imports
func (s *IdeasService) Import(ideasImportAttributes IdeasImportAttributes) (*IdeasImportResponse, error) {
	path := "/api/ideas/actions/import"
	ideasImportResponse := &IdeasImportResponse{}

	resp, err := s.client.post(path, ideasImportAttributes, ideasImportResponse)

	// import does not return any response data, so expect the response decode to be EOF
	if err != nil && err != io.EOF {
		return nil, err
	}

	ideasImportResponse.HTTPResponse = resp
	return ideasImportResponse, nil
}
