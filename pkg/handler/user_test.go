package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"chitchat/pkg/model"
)

func TestGetAllUsers(t *testing.T) {

	request, _ := http.NewRequest(http.MethodGet, "/api/v1/users", nil)
	response := httptest.NewRecorder()

	GetAllUsers(response, request)

	got := response.Body
	var want []model.User

	err := json.NewDecoder(got).Decode(&want)
	if err != nil {
		t.Fatalf("Response %v doesn't match User structure %v", got, err)
	}
}

func TestGetUserById(t *testing.T) {

	t.Run("get existing user", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/api/v1/users/567", nil)
		response := httptest.NewRecorder()

		GetUserById(response, request)

		got := response.Body
		var want model.User

		err := json.NewDecoder(got).Decode(&want)
		if err != nil {
			t.Fatalf("Response %v doesn't match user structure %v", got, err)
		}
	})

	t.Run("get non-existing user", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/api/v1/users/941", nil)
		response := httptest.NewRecorder()

		GetUserById(response, request)

		got := response.Code
		var want int = 404

		if got != want {
			t.Fatal("User should not exist")
		}
	})

}
