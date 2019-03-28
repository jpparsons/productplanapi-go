package productplan

import (
	"io"
	"net/http"
	"testing"
	"time"
)

func TestIdeas_Import(t *testing.T) {
	setupMockServer()
	defer teardownMockServer()

	mux.HandleFunc("/api/ideas/actions/import", func(w http.ResponseWriter, r *http.Request) {
		httpResponse := httpResponseFixture(t, "/ideas_import/idea_import_success.http")

		testMethod(t, r, "POST")
		testHeaders(t, r)

		w.WriteHeader(httpResponse.StatusCode)
		io.Copy(w, httpResponse.Body)
	})

	roadmapID := 4946

	d := []time.Duration{time.Second}
	timestamps := Timestamps{CreatedAt: time.Date(2018, 3, 31, 14, 36, 40, 07, time.Local).Round(d[0]),
		UpdatedAt: time.Date(2018, 3, 31, 15, 17, 52, 07, time.Local).Round(d[0])}

	links := IdeaLinks{Roadmap: map[string]string{"href": "/api/roadmaps/4946"},
		ExternalLinks: map[string]string{"href": "/api/ideas/110689/external_links"}}

	idea := Ideas{
		Href:           "/api/ideas/110689",
		ID:             110689,
		Name:           "Product Research",
		Description:    "Research and test models",
		StrategicValue: "Leads leads to a better product",
		Notes:          "Engage customers as early adopters",
		PercentDone:    10,
		Effort:         2,
		Tags:           []string{"devops", "security"},
		Fields:         map[string]string{"pp_lanes": "Lane 2", "pp_legend": "Goal 1"},
		Timestamps:     &timestamps,
		IdeaLinks:      &links,
	}

	ideas := IdeasImportAttributes{
		IdeaImportRoadmap: IdeaImportRoadmap{ID: roadmapID},
		Ideas:             []Ideas{idea},
	}

	ideasResponse, err := client.Ideas.Import(ideas)
	if err != nil {
		t.Fatalf("Ideas.Import(ideas) returned error: %v", err)
	}

	response := ideasResponse
	if response.HTTPResponse.StatusCode != 202 {
		t.Errorf("response.HTTPResponse.StatusCode GOT: %+v", response.HTTPResponse.StatusCode)
	}

}
