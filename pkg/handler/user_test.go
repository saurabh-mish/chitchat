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
		t.Errorf("Response %v doesn't match User structure %v", got, err)
	}
}
