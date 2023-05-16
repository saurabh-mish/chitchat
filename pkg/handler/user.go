package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"chitchat/pkg/model"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := []model.User{
		{
			Id:       1234,
			Uuid:     "long-string-1",
			Name:     "John",
			Email:    "john@gmail.com",
			Password: "passw0rd",
		},
		{
			Id:       987,
			Uuid:     "long-string-2",
			Name:     "Jane",
			Email:    "jane@outlook.com",
			Password: "pa$$word",
		},
	}

	//allUsers, _ := json.Marshal(users)
	//fmt.Fprintf(writer, string(allUsers))
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		fmt.Fprintf(w, "Unable to parse defined data to User structure...")
	}
}
