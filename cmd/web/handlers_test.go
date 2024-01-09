package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApplicationHandlers(t *testing.T) {
	var theTests = []struct {
		name               string
		url                string
		expectedStatusCode int
	}{
		{"home", "/", http.StatusOK},
	}

	var app application
	routes := app.routes()

	ts := httptest.NewTLSServer(routes)
	defer ts.Close()
	pathToTemplates = "./../../templates/"

	for _, e := range theTests {
		res, err := ts.Client().Get(ts.URL + e.url)
		if err != nil {
			t.Log(err)
			t.Fail()
		}

		if res.StatusCode != e.expectedStatusCode {
			t.Errorf("Got %d, expected %d for %s", res.StatusCode, e.expectedStatusCode, e.name)
		}
	}
}
