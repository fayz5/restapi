package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/fayz5/restapi/handlers"
)

func main() {
	//Init Mux
	r := mux.NewRouter()
	//Register Handlers
	r.HandleFunc("/api/message", handlers.MessageHandler)

	fmt.Println("Listening on Port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

