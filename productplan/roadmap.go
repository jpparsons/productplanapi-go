package productplan

// Roadmap represents a roadmap name
type Roadmap struct {
	Href         string   `json:"href,omitempty"`
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Description  string   `json:"description,omitempty"`
	OwnerEmail   string   `json:"owner_email,omitempty"`
	IsVersion    bool     `json:"is_version,omitempty"`
	CopiedFrom   *Roadmap `json:"copied_from,omitempty"`
	Timestamps   `json:"timestamps"`
	RoadmapLinks `json:"links,omitempty"`
}

// RoadmapLinks on an idea
type RoadmapLinks struct {
	Bars         map[string]string `json:"bars,omitempty"`
	Ideas        map[string]string `json:"ideas,omitempty"`
	CustomFields map[string]string `json:"custom_fields,omitempty"`
}

// RoadmapListOptions specifies optional parameters to pass to Roadmaps.ListRoadmaps method
type RoadmapListOptions struct {
	// option to return roadmaps that are shared
	IncludeShared bool `url:"include_shared,omitempty"`

	// option to return roadmap versions
	IncludeVersions bool `url:"include_versions,omitempty"`

	ListOptions
}
