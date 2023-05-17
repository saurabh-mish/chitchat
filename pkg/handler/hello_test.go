package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {

	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	Hello(response, request)

	if response.Body.String() != "Hello, World!\n" {
		t.Errorf("Invalid response body from server; got %v, want %v", got, want)
	}

	if response.Code != 200 {
		t.Errorf("Invalid response status code from server; got %v, want %v", got, want)
	}
}
