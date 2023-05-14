package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloServer(t *testing.T) {
	t.Run("valid response from server", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		Hello(response, request)

		got := response.Body.String()
		want := "Hello, World!\n"

		if got != want {
			t.Errorf("Invalid response from server; got %q, want %q", got, want)
		}
	})
}
