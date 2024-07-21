package main

import (
	"net/http"
	"api_20240721/handler"
)

func main() {
	http.HandleFunc("/", handler.Router)
	http.ListenAndServe(":8000", nil)
}
