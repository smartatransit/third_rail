package controllers

import (
	"encoding/json"
	"github.com/smartatransit/third_rail/pkg/validators"
	"net/http"
)

func GetStaticScheduleByStation(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response{})
}

func GetLines(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	mev := validators.NewMartaEntitiesValidator()
	lines, _ := mev.GetEntities(validators.MARTA_LINES)

	response := linesResponse{linesData{Lines: lines}}

	json.NewEncoder(w).Encode(response)
}

func GetDirections(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	mev := validators.NewMartaEntitiesValidator()
	directions, _ := mev.GetEntities(validators.MARTA_DIRECTIONS)

	response := directionsResponse{directionsData{Directions: directions}}

	json.NewEncoder(w).Encode(response)
}

func GetStations(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	mev := validators.NewMartaEntitiesValidator()
	stations, _ := mev.GetEntities(validators.MARTA_STATIONS)

	response := stationsResponse{stationsData{Stations: stations}}

	json.NewEncoder(w).Encode(response)
}

func GetLocations(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	mev := validators.NewMartaEntitiesValidator()
	stations, _ := mev.GetEntities(validators.MARTA_STATIONS)

	response := stationsResponse{stationsData{Stations: stations}}

	json.NewEncoder(w).Encode(response)
}
