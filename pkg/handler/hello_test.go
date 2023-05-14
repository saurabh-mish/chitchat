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

	t.Run("response body", func(t *testing.T) {
		got := response.Body.String()
		want := "Hello, World!\n"

		if got != want {
			t.Errorf("Invalid response body from server; got %v, want %v", got, want)
		}
	})

	t.Run("response status code", func(t *testing.T) {
		got := response.Code
		want := 200

		if got != want {
			t.Errorf("Invalid response status code from server; got %v, want %v", got, want)
		}
	})
}
