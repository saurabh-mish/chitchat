package main

import (
	"log"
	"net/http"

	"chitchat/pkg/handler"
)

func main() {
	http.HandleFunc("/hello", handler.Hello)
	http.HandleFunc("/api/v1/users", handler.GetAllUsers)
	http.HandleFunc("/api/v1/users/", handler.GetUserById)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
