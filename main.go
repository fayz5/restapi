package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fayz5/restapi/handlers"
	"github.com/gorilla/mux"
)

func main() {
	//Init Mux
	r := mux.NewRouter()
	//Register Handlers
	r.HandleFunc("/api/message", handlers.MessageHandler)
	//Log server status
	fmt.Println("Listening on Port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
