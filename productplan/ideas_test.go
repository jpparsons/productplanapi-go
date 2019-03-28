package productplan

import (
	"io"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestIdeas_Show(t *testing.T) {
	setupMockServer()
	defer teardownMockServer()

	mux.HandleFunc("/api/ideas/110689", func(w http.ResponseWriter, r *http.Request) {
		httpResponse := httpResponseFixture(t, "/ideas/success.http")

		testMethod(t, r, "GET")
		testHeaders(t, r)

		w.WriteHeader(httpResponse.StatusCode)
		io.Copy(w, httpResponse.Body)
	})

	ideaID := "110689"

	ideasResponse, err := client.Ideas.Show(ideaID)
	if err != nil {
		t.Fatalf("Ideas.Show(ideaID) returned error: %v", err)
	}

	//fmt.Println(ideasResponse.Ideas.Timestamps.CreatedAt)
	//fmt.Println(ideasResponse.Ideas.IdeaLinks.Roadmap["href"])
	//fmt.Println("type of: ", reflect.TypeOf(ideasResponse.Ideas.Timestamps.CreatedAt))

	d := []time.Duration{time.Second}
	timestamps := Timestamps{CreatedAt: time.Date(2017, 3, 31, 14, 36, 40, 07, time.Local).Round(d[0]),
		UpdatedAt: time.Date(2017, 3, 31, 15, 17, 52, 07, time.Local).Round(d[0])}

	links := IdeaLinks{Roadmap: map[string]string{"href": "/api/roadmaps/5912"},
		ExternalLinks: map[string]string{"href": "/api/ideas/110689/external_links"}}

	want := Ideas{
		Href:           "/api/ideas/110689",
		ID:             110689,
		Name:           "ImplementationIdeaTest",
		Description:    "Implement and test the feature.",
		StrategicValue: "polarity",
		Notes:          "autotoxication",
		PercentDone:    78,
		Effort:         2,
		Tags:           []string{"electrotypy", "monoservice"},
		Fields:         map[string]string{"pp_lanes": "Lane 2", "pp_legend": "Goal 1"},
		Timestamps:     &timestamps,
		IdeaLinks:      &links,
	}

	got := ideasResponse.Ideas
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ideasResponse.Ideas() returned GOT: %+v, WANT %+v", got, want)
	}

}
