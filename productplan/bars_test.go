package productplan

import (
	"io"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestBarsService_ListBars_WithOptions(t *testing.T) {
	setupMockServer()
	defer teardownMockServer()

	mux.HandleFunc("/api/bars", func(w http.ResponseWriter, r *http.Request) {
		httpResponse := httpResponseFixture(t, "/bars/list_bars_with_name_filter.http")

		testMethod(t, r, "GET")
		testHeaders(t, r)

		w.WriteHeader(httpResponse.StatusCode)
		io.Copy(w, httpResponse.Body)
	})

	barName := "APIBar1001"

	barsResponse, err := client.Bars.ListBars(&ListOptions{Filters: "name=" + barName})
	if err != nil {
		t.Fatalf("Bars.ListBars() returned error: %v", err)
	}

	d := []time.Duration{time.Second}
	timestamps := Timestamps{CreatedAt: time.Date(2017, 11, 27, 16, 51, 37, 07, time.Local).Round(d[0]),
		UpdatedAt: time.Date(2018, 07, 11, 07, 51, 05, 07, time.Local).Round(d[0])}

	barLinks := BarLinks{
		Roadmap:       map[string]string{"href": "/api/roadmaps/4946"},
		ChildBars:     map[string]string{"href": "/api/bars/205414/child_bars"},
		ExternalLinks: map[string]string{"href": "/api/bars/205414/external_links"}}

	want := Bar{
		Href:           "/api/bars/205414",
		ID:             205414,
		Name:           "APIBar1001",
		StartDate:      "2017-01-21",
		EndDate:        "2017-04-19",
		Description:    "bar desc",
		StrategicValue: "low",
		Notes:          "notes added",
		PercentDone:    0,
		Effort:         1,
		Tags:           []string{"tag1"},
		Fields:         map[string]string{"pp_lanes": "Lane 2", "pp_legend": "Goal 4"},
		Timestamps:     timestamps,
		BarLinks:       barLinks,
	}

	bar := (*barsResponse)[0]
	got := bar.Bar

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Bar.ListBars returned GOT: %+v, WANT %+v", got, want)
	}

}

func TestBarsService_UpdateBar(t *testing.T) {
	setupMockServer()
	defer teardownMockServer()

	mux.HandleFunc("/api/bars/205400", func(w http.ResponseWriter, r *http.Request) {
		httpResponse := httpResponseFixture(t, "/bars/bar_update.http")

		testMethod(t, r, "PATCH")
		testHeaders(t, r)

		w.WriteHeader(httpResponse.StatusCode)
		io.Copy(w, httpResponse.Body)
	})

	barAttributes := UpdateBar{
		StartDate: "2018-01-07",
		EndDate:   "2018-04-19",
	}

	barsResponse, err := client.Bars.UpdateBar(205400, barAttributes)
	if err != nil {
		t.Fatalf("Bars.UpdateBar(bar) returned error: %v", err)
	}

	if barsResponse.HTTPResponse.StatusCode != 204 {
		t.Errorf("barsResponse.HTTPResponse.StatusCode GOT: %+v", barsResponse.HTTPResponse.StatusCode)
	}
}
