package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type responseMessage struct {
	Status	string `json:"status"`
	Massage	string `json:"message"`
	Timestamp	string `json:"timestamp"`
}

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", hello).Methods("GET")
	return r
}

func hello(w http.ResponseWriter, r *http.Request) {
	log.Printf(">>> GET")
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	resp := responseMessage{Status: "ok", Massage: "Hello World !!!", Timestamp: currentTime}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
	log.Printf("<<<(%d)%v", http.StatusOK, resp)
	return
}
