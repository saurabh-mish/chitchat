package handler

import (
	"fmt"
	"net/http"
)

func Hello(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello, World!\n")
}
