package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestApplicationAddIpToContext(t *testing.T) {
	tests := []struct {
		headerName  string
		headerValue string
		address     string
		emptyAddr   bool
	}{
		{"", "", "", false},
		{"", "", "", true},
		{"X-Forwarded-For", "192.3.2.1", "", false},
		{"", "", "hello:world", false},
	}

	nextHandler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		val := request.Context().Value(contextUserKey)
		if val == nil {
			t.Error(contextUserKey, "not present")
		}
		_, ok := val.(string)
		if !ok {
			t.Error("not string")
		}
	})
	for _, e := range tests {
		handlerToTest := app.addIpToContext(nextHandler)

		req := httptest.NewRequest("GET", "http://testing", nil)
		if e.emptyAddr {
			req.RemoteAddr = ""
		}

		if len(e.headerName) > 0 {
			req.Header.Add(e.headerName, e.headerValue)
		}

		if len(e.address) > 0 {
			req.RemoteAddr = e.address
		}
		handlerToTest.ServeHTTP(httptest.NewRecorder(), req)
	}
}
func TestApplicationIpFromContext(t *testing.T) {

	ctx := context.Background()
	ctx = context.WithValue(ctx, contextUserKey, "tjosan")
	ip := app.ipFromContext(ctx)

	if !strings.EqualFold("tjosan", ip) {
		t.Error("wrong value returned: ", ip)
	}
}
