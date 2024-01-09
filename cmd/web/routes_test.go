package main

import (
	"github.com/go-chi/chi"
	"net/http"
	"strings"
	"testing"
)

func TestApplicationRoutes(t *testing.T) {
	var registred = []struct {
		route  string
		method string
	}{
		{"/", "GET"},
		{"/login", "POST"},
		{"/static/*", "GET"},
	}

	mux := app.routes()
	chiRoutes := mux.(chi.Routes)

	for _, route := range registred {
		if !routeExists(route.route, route.method, chiRoutes) {
			t.Errorf("%s, is not registred", route.route)
		}
	}
}

func routeExists(testRoute, testMethod string, chiRoutes chi.Routes) bool {
	found := false

	_ = chi.Walk(chiRoutes, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		if strings.EqualFold(method, testMethod) && strings.EqualFold(route, testRoute) {
			found = true
		}
		return nil
	})
	return found
}
