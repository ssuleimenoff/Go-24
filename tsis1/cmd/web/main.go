// main.go
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/api/heroes", GetHeroes).Methods("GET")
	router.HandleFunc("/api/heroes/{name}", GetHero).Methods("GET")
	router.HandleFunc("/health", HealthCheck).Methods("GET")

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", router))
}
