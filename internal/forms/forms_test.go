package forms

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()

	if !isValid {
		t.Error("got invalid when it should be valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")

	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}

	postedData := url.Values{}

	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	r = httptest.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData

	form = New(r.PostForm)

	form.Required("a", "b", "c")

	if !form.Valid() {
		t.Error("form shows NOT valid when required fields are NOT missing")
	}
}

func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	if form.Has("test") {
		t.Error("Form type should not have `test` field.")
	}

	postedData := url.Values{}

	postedData.Add("b", "a")

	r = httptest.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData

	form = New(r.PostForm)

	if !form.Has("b") {
		t.Error("Form should have field `b`, but this might be due to the request object not having the form at the start.")
	}
}

func TestForm_MinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	postedData := url.Values{}

	postedData.Add("shortTest", "a")
	r.PostForm = postedData

	form := New(r.PostForm)
	form.MinLength("shortTest", 2)

	if form.Valid() {
		t.Error("Form should not be valid, min length for `shortTest` is 2.")
	}

	isError := form.Errors.Get("shortTest")
	if isError == "" {
		t.Error("Should have an error, but did not get one")
	}

	r = httptest.NewRequest("POST", "/whatever", nil)
	postedData = url.Values{}

	postedData.Add("longTest", "aaaa")
	r.PostForm = postedData

	form = New(r.PostForm)
	form.MinLength("longTest", 2)

	if !form.Valid() {
		t.Error("Form should BE valid, min length for `longTest` is 2, we have 4.")
	}

	isError = form.Errors.Get("longTest")
	if isError != "" {
		t.Error("Should NOT have an error, but got one")
	}
}

func TestForm_IsEmail(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	postedData := url.Values{}

	postedData.Add("badEmail", "a")
	r.PostForm = postedData

	form := New(r.PostForm)
	form.IsEmail("badEmail")

	if form.Valid() {
		t.Error("Form should not be valid, not an email")
	}

	r = httptest.NewRequest("POST", "/whatever", nil)
	postedData = url.Values{}

	postedData.Add("goodEmail", "a@example.com")
	r.PostForm = postedData

	form = New(r.PostForm)
	form.IsEmail("goodEmail")

	if !form.Valid() {
		t.Error("Form should BE valid, field has an example of a good email.")
	}
}
