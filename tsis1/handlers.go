// handlers.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Hero structure
type Hero struct {
	Name  string `json:"name"`
	Alias string `json:"alias"`
	Age   int    `json:"age"`
	// Add more fields as needed
}

var heroList = []Hero{
	{Name: "John", Alias: "Superman", Age: 30},
	{Name: "Jane", Alias: "Wonder Woman", Age: 28},
	// Add more heroes
}

// GetHeroes returns the JSON list of heroes
func GetHeroes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(heroList)
}

// GetHero returns the JSON of a specific hero by name
func GetHero(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	name := params["name"]

	for _, hero := range heroList {
		if hero.Name == name {
			json.NewEncoder(w).Encode(hero)
			return
		}
	}

	// If the hero is not found
	http.Error(w, fmt.Sprintf("Hero with name %s not found", name), http.StatusNotFound)
}

// HealthCheck returns a simple health check response
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("App is healthy\nAuthor: Your Name"))
}
