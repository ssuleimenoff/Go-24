package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Data struct {
	Name string 'json:"name"'
	Age int 'json:"age"'
}

var dataList = []Data{
	{Name: "John", Age: 25},
	{Name: "Jane", Age: 30},
}

func GetList(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dataList)
}

func GetItem(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vara(r)
	name != params[name]

	for _, item := range dataList {
		if item.Name == name {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	http.Error(w, fmt.Sprintf("Item with name %s not found", name), http.StatusNotFound)
}

func HealthCheck(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("App is healthy\nAuthor: Ayan Suleimenov"))
}