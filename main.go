package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Message represents data model for JSON api
type Message struct {
	Text      string `json:"text"`
	Timestamp int64  `json:"timestamp"`
}

func main() {
	//Init Mux
	r := mux.NewRouter()
	//Register Handlers
	r.HandleFunc("/api/message", messageHandler)

	fmt.Println("Listening on Port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	m := Message{"Hello From Server", time.Now().Unix()}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(m)
}
