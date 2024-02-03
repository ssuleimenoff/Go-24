package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ssuleimenoff/Go-24/tsis1/api"
)

func GetBands(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(api.Bands)
}
func GetBand(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)

	for _, item := range api.Bands {
		if item.Id == params["id"] {
			json.NewEncoder(writer).Encode(item)
			return
		}
	}
	json.NewEncoder(writer).Encode(&api.Band{})
}
func UpdateBand(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)

	for index, item := range api.Bands {
		if item.Id == params["id"] {
			api.Bands = append(api.Bands[:index], api.Bands[index+1:]...)

			var band api.Band
			_ = json.NewDecoder(request.Body).Decode(&band)

			band.Id = strconv.Itoa(rand.Intn(1000000))

			api.Bands = append(api.Bands, band)
			json.NewEncoder(writer).Encode(band)
			return
		}

	}
	json.NewEncoder(writer).Encode(api.Bands)
}
func CreateBand(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var band api.Band
	_ = json.NewDecoder(request.Body).Decode(&band)

	band.Id = strconv.Itoa(rand.Intn(1000000))

	api.Bands = append(api.Bands, band)
	json.NewEncoder(writer).Encode(band)
}
func DeleteBand(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)
	for index, item := range api.Bands {
		if item.Id == params["id"] {
			api.Bands = append(api.Bands[:index], api.Bands[index+1:]...)
			break
		}
	}
	json.NewEncoder(writer).Encode(api.Bands)
}
