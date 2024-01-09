package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Has(t *testing.T) {
	form := NewForm(nil)
	has := form.Has("something")
	if has {
		t.Error("Form has field when it should not")
	}
	postedData := url.Values{}
	postedData.Add("a", "a")

	form = NewForm(postedData)
	has = form.Has("a")
	if !has {
		t.Error("Form does not have value, when it should")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/something", nil)
	form := NewForm(r.PostForm)
	form.Required("a", "b", "c")

	if form.Valid() {
		t.Error("form is valid when values are missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "b")
	postedData.Add("c", "c")

	r, _ = http.NewRequest("POST", "/somewhare", nil)
	r.PostForm = postedData

	form = NewForm(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("valid form displays invalid")
	}
}

func TestForm_Check(t *testing.T) {
	form := NewForm(nil)
	form.Check(false, "password", "password is required")
	if form.Valid() {
		t.Error("Displays valid, when it's not")
	}
}

func TestForm_Get(t *testing.T) {
	form := NewForm(nil)
	form.Check(false, "password", "password is required")

	s := form.Errors.Get("password")
	if len(s) == 0 {
		t.Error("did not display error message")
	}

	s = form.Errors.Get("something")
	if len(s) != 0 {
		t.Error("did display error, when no error was expected")
	}

}
