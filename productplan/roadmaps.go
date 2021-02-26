package productplan

import (
	"fmt"
)

// RoadmapsService handles communication with the roadmap
// methods of the Productplan API.
type RoadmapsService struct {
	client *Client
}

// RoadmapsResponse represents a response from an API method that returns a Roadmap struct.
type RoadmapsResponse struct {
	Roadmap
}

// ListRoadmaps get a list of roadmaps
func (s *RoadmapsService) ListRoadmaps(options *RoadmapListOptions) (*[]RoadmapsResponse, error) {
	path := "/api/roadmaps"
	var roadmapsResponse *[]RoadmapsResponse

	path, err := addURLQueryOptions(path, options)
	if err != nil {
		return nil, err
	}

	_, err = s.client.get(path, &roadmapsResponse)
	if err != nil {
		return nil, err
	}

	return roadmapsResponse, nil
}

// GetRoadmap roadmap by ID
func (s *RoadmapsService) GetRoadmap(id int) (*RoadmapsResponse, error) {
	path := fmt.Sprintf("/api/roadmaps/%v", id)
	var roadmapsResponse *RoadmapsResponse

	_, err := s.client.get(path, &roadmapsResponse)
	if err != nil {
		return nil, err
	}

	return roadmapsResponse, nil
}

// GetBars get bars on a roadmap
func (s *RoadmapsService) GetBars(roadmap Roadmap) (*[]BarsResponse, error) {
	path := fmt.Sprintf("/api/roadmaps/%v/bars", roadmap.ID)
	// https://github.com/DaveAppleton/LoadObjectSliceFromJson
	var barsResponse *[]BarsResponse

	_, err := s.client.get(path, &barsResponse)
	if err != nil {
		return nil, err
	}

	return barsResponse, nil
}
