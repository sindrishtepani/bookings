package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNoSurve(t *testing.T) {
	var testHandler myHandler

	h := NoSurf(&testHandler)

	switch v := h.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Error(fmt.Printf("type is not http.Handler, but is %T", v))
	}
}

func TestSessionLoad(t *testing.T) {
	var testHandler myHandler

	h := SessionLoad(&testHandler)

	switch v := h.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Error(fmt.Printf("type is not http.Handler, but is %T", v))
	}
}
