package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Routes Defining
	router.HandleFunc("api/data", GetList).Methods("GET")
	router.HandleFunc("api/data/{name}, GetItem").Methods("GET")
	router.HandleFunc("/health", HealthCheck).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
