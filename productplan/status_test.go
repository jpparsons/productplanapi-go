package productplan

import (
	"io"
	"net/http"
	"reflect"
	"testing"
)

func TestStatus(t *testing.T) {
	setupMockServer()
	defer teardownMockServer()

	mux.HandleFunc("/api/status", func(w http.ResponseWriter, r *http.Request) {
		httpResponse := httpResponseFixture(t, "/status/success.http")

		testMethod(t, r, "GET")
		testHeaders(t, r)

		w.WriteHeader(httpResponse.StatusCode)
		io.Copy(w, httpResponse.Body)
	})

	statusResponse, err := client.Status.GetStatus()
	if err != nil {
		t.Fatalf("Status.GetStatus() returned error: %v", err)
	}

	/*appStatus := statusResponse.Status.Application
	if appStatus != "up" {
		t.Errorf("statusResponse.Status.Application returned %+v", appStatus)
	}
	dbStatus := statusResponse.Status.Database
	if dbStatus != "up" {
		t.Errorf("statusResponse.Status.Database returned %+v", dbStatus)
	}*/

	status := statusResponse.Status
	want1 := Status{Application: "up", Database: "up"}
	if !reflect.DeepEqual(status, want1) {
		t.Errorf("statusResponse.Status() returned %+v, want %+v", status, want1)
	}

	meta := statusResponse.Metadata
	want2 := Metadata{Application: "ProductPlan API", Version: 1}
	if !reflect.DeepEqual(meta, want2) {
		t.Errorf("statusResponse.Status() returned %+v, want %+v", meta, want2)
	}
}
