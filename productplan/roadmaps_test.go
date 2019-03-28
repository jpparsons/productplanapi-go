package productplan

import (
	"fmt"
	"io"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestRoadmapsService_ListRoadmaps_WithOptions(t *testing.T) {
	setupMockServer()
	defer teardownMockServer()

	mux.HandleFunc("/api/roadmaps", func(w http.ResponseWriter, r *http.Request) {
		httpResponse := httpResponseFixture(t, "/roadmaps/list_roadmaps_success.http")

		testMethod(t, r, "GET")
		testHeaders(t, r)

		w.WriteHeader(httpResponse.StatusCode)
		io.Copy(w, httpResponse.Body)
	})

	roadmapName := "Roadmap10"

	roadmapsResponse, err := client.Roadmaps.ListRoadmaps(&ListOptions{Filters: "name=" + roadmapName})
	if err != nil {
		t.Fatalf("Roadmaps.ListRoadmaps() returned error: %v", err)
	}

	roadmap0 := (*roadmapsResponse)[0]

	d := []time.Duration{time.Second}
	timestamps := Timestamps{CreatedAt: time.Date(2017, 10, 03, 12, 58, 07, 07, time.Local).Round(d[0]),
		UpdatedAt: time.Date(2017, 10, 05, 12, 02, 07, 07, time.Local).Round(d[0])}

	roadmapLinks := RoadmapLinks{
		Bars:         map[string]string{"href": "/api/roadmaps/7302/bars"},
		Ideas:        map[string]string{"href": "/api/roadmaps/7302/ideas"},
		CustomFields: map[string]string{"href": "/api/roadmaps/7302/custom_fields"}}

	want0 := Roadmap{
		Href:         "/api/roadmaps/7302",
		ID:           7302,
		Name:         "Roadmap10",
		Description:  "Load testing plan",
		OwnerEmail:   "user@testmail.com",
		Timestamps:   timestamps,
		RoadmapLinks: roadmapLinks,
	}

	got0 := roadmap0.Roadmap
	if !reflect.DeepEqual(got0, want0) {
		t.Errorf("roadmap.Roadmap returned GOT: %+v, WANT %+v", got0, want0)
	}

	// TODO: CHANGE roadmap ID
	roadmap1 := (*roadmapsResponse)[1]
	want1 := Roadmap{
		Href:         "/api/roadmaps/7302",
		ID:           7302,
		Name:         "Roadmap11",
		Description:  "Load testing plan",
		OwnerEmail:   "user@testmail.com",
		Timestamps:   timestamps,
		RoadmapLinks: roadmapLinks,
	}

	got1 := roadmap1.Roadmap
	if !reflect.DeepEqual(got1, want1) {
		t.Errorf("roadmap.Roadmap returned GOT: %+v, WANT %+v", got1, want1)
	}

}

func TestRoadmapsService_ListRoadmaps_WithOptions_NotFound(t *testing.T) {
	setupMockServer()
	defer teardownMockServer()

	mux.HandleFunc("/api/roadmaps", func(w http.ResponseWriter, r *http.Request) {
		httpResponse := httpResponseFixture(t, "/roadmaps/list_roadmaps_not_found.http")

		testMethod(t, r, "GET")
		testHeaders(t, r)

		w.WriteHeader(httpResponse.StatusCode)
		io.Copy(w, httpResponse.Body)
	})

	roadmapName := "NotFound"

	roadmapsResponse, err := client.Roadmaps.ListRoadmaps(&ListOptions{Filters: "name=" + roadmapName})
	if err != nil {
		t.Fatalf("Roadmaps.ListRoadmaps() returned error: %v", err)
	}

	got := *roadmapsResponse
	want := []RoadmapsResponse{}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Roadmaps.ListRoadmaps returned GOT: %+v, WANT %+v", got, want)
	}

}

func TestRoadmapsService_ListRoadmaps(t *testing.T) {
	setupMockServer()
	defer teardownMockServer()

	mux.HandleFunc("/api/roadmaps", func(w http.ResponseWriter, r *http.Request) {
		httpResponse := httpResponseFixture(t, "/roadmaps/list_roadmaps_success.http")

		testMethod(t, r, "GET")
		testHeaders(t, r)

		w.WriteHeader(httpResponse.StatusCode)
		io.Copy(w, httpResponse.Body)
	})

	roadmapsResponse, err := client.Roadmaps.ListRoadmaps(nil)
	if err != nil {
		t.Fatalf("Roadmaps.ListRoadmaps() returned error: %v", err)
	}

	for _, roadmap := range *roadmapsResponse {
		fmt.Println(roadmap.Name)
	}

}

func TestRoadmapsService_GetRoadmap(t *testing.T) {
	setupMockServer()
	defer teardownMockServer()

	mux.HandleFunc("/api/roadmaps/7302", func(w http.ResponseWriter, r *http.Request) {
		httpResponse := httpResponseFixture(t, "/roadmaps/get_roadmap_success.http")

		testMethod(t, r, "GET")
		testHeaders(t, r)

		w.WriteHeader(httpResponse.StatusCode)
		io.Copy(w, httpResponse.Body)
	})

	roadmapsResponse, err := client.Roadmaps.GetRoadmap(7302)
	if err != nil {
		t.Fatalf("Roadmaps.ListRoadmaps() returned error: %v", err)
	}

	d := []time.Duration{time.Second}
	timestamps := Timestamps{CreatedAt: time.Date(2017, 06, 01, 12, 03, 53, 07, time.Local).Round(d[0]),
		UpdatedAt: time.Date(2018, 06, 27, 15, 52, 32, 07, time.Local).Round(d[0])}

	roadmapLinks := RoadmapLinks{
		Bars:         map[string]string{"href": "/api/roadmaps/7302/bars"},
		Ideas:        map[string]string{"href": "/api/roadmaps/7302/ideas"},
		CustomFields: map[string]string{"href": "/api/roadmaps/7302/custom_fields"}}

	want := Roadmap{
		Href:         "/api/roadmaps/7302",
		ID:           7302,
		Name:         "Roadmap10",
		Description:  "roadmap description",
		OwnerEmail:   "user@testmail.com",
		Timestamps:   timestamps,
		RoadmapLinks: roadmapLinks,
	}

	got := roadmapsResponse.Roadmap
	if !reflect.DeepEqual(got, want) {
		t.Errorf("roadmap.Roadmap returned GOT: %+v, WANT %+v", got, want)
	}
}

func TestRoadmapsService_GetBars(t *testing.T) {
	setupMockServer()
	defer teardownMockServer()

	mux.HandleFunc("/api/roadmaps/7302/bars", func(w http.ResponseWriter, r *http.Request) {
		httpResponse := httpResponseFixture(t, "/roadmaps/get_bars_success.http")

		testMethod(t, r, "GET")
		testHeaders(t, r)

		w.WriteHeader(httpResponse.StatusCode)
		io.Copy(w, httpResponse.Body)
	})

	d := []time.Duration{time.Second}
	timestamps := Timestamps{CreatedAt: time.Date(2017, 10, 03, 12, 58, 07, 07, time.Local).Round(d[0]),
		UpdatedAt: time.Date(2017, 10, 05, 12, 02, 07, 07, time.Local).Round(d[0])}

	roadmapLinks := RoadmapLinks{
		Bars:         map[string]string{"href": "/api/roadmaps/7302/bars"},
		Ideas:        map[string]string{"href": "/api/roadmaps/7302/ideas"},
		CustomFields: map[string]string{"href": "/api/roadmaps/7302/custom_fields"}}

	roadmap := Roadmap{
		Href:         "/api/roadmaps/7302",
		ID:           7302,
		Name:         "Roadmap10",
		Description:  "Load testing plan",
		OwnerEmail:   "user@testmail.com",
		Timestamps:   timestamps,
		RoadmapLinks: roadmapLinks,
	}

	roadmapsResponse, err := client.Roadmaps.GetBars(roadmap)
	if err != nil {
		t.Fatalf("Roadmaps.GetBars() returned error: %v", err)
	}

	barLinks := BarLinks{
		Roadmap:       map[string]string{"href": "/api/roadmaps/7302"},
		ChildBars:     map[string]string{"href": "/api/bars/1102402/child_bars"},
		ExternalLinks: map[string]string{"href": "/api/bars/1102402/external_links"}}

	want0 := Bar{
		Href:           "/api/bars/110240",
		ID:             110240,
		Name:           "API Bar",
		StartDate:      "2017-06-21",
		EndDate:        "2017-09-21",
		Description:    "desc",
		StrategicValue: "low",
		Notes:          "notes",
		PercentDone:    0,
		Effort:         5,
		Tags:           []string{"ssl", "docker"},
		Fields:         map[string]string{"pp_lanes": "Lane 2", "pp_legend": "Goal 4"},
		Timestamps:     timestamps,
		BarLinks:       barLinks,
	}

	bar0 := (*roadmapsResponse)[0]
	got0 := bar0.Bar

	if !reflect.DeepEqual(got0, want0) {
		t.Errorf("Roadmaps.GetBars returned GOT: %+v, WANT %+v", got0, want0)
	}

}
