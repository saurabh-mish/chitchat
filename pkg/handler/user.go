package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

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

func GetUserById(w http.ResponseWriter, r *http.Request) {

	user := model.User{
		Id:       567,
		Uuid:     "long-string-3",
		Name:     "Bob",
		Email:    "bob@gmail.com",
		Password: "decentP@ssw0rd",
	}

	id := strings.TrimPrefix(r.URL.Path, "/api/v1/users/")

	if id != strconv.Itoa(user.Id) {
		//fmt.Fprintf(w, "User does not exist...")
		w.WriteHeader(http.StatusNotFound)
	} else {
		err := json.NewEncoder(w).Encode(user)
		if err != nil {
			fmt.Fprintf(w, "Unable to parse defined structure to user structure ...")
		}
	}
}
