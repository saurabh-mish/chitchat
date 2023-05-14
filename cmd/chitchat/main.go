package main

import (
	"log"
	"net/http"

	"chitchat/pkg/handler"
)

func main() {
	http.HandleFunc("/hello", handler.Hello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
