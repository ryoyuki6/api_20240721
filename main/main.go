package main

import (
	"net/http"
	"log"
	"api_20240721/handler"
)

func main() {
	log.Printf("api_20240721 - 20240723")
	r := handler.Router()
	http.Handle("/", r)
	port := "8000"
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Printf("err %s", err.Error())
	}
}
