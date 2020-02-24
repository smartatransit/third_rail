package controllers

import (
	"encoding/json"
	"github.com/smartatransit/third_rail/pkg/transformers"
	"github.com/smartatransit/third_rail/pkg/validators"
	"net/http"
	"strconv"
)

type StaticController struct {
}

func (controller StaticController) GetStaticScheduleByStation(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	v := req.URL.Query()

	schedule := v.Get("schedule")
	stationName := v.Get("station_name")

	if len(schedule) < 1 || len(stationName) < 1 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	json.NewEncoder(w).Encode(response{})
}

func (controller StaticController) GetLines(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	mev := validators.NewMartaEntitiesValidator()
	lines, _ := mev.GetEntities(validators.MARTA_LINES)

	response := linesResponse{linesData{Lines: lines}}

	json.NewEncoder(w).Encode(response)
}

func (controller StaticController) GetDirections(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	mev := validators.NewMartaEntitiesValidator()
	directions, _ := mev.GetEntities(validators.MARTA_DIRECTIONS)

	response := directionsResponse{directionsData{Directions: directions}}

	json.NewEncoder(w).Encode(response)
}

func (controller StaticController) GetStations(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	mev := validators.NewMartaEntitiesValidator()
	stations, _ := mev.GetEntities(validators.MARTA_STATIONS)

	response := stationsResponse{stationsData{Stations: stations}}

	json.NewEncoder(w).Encode(response)
}

func (controller StaticController) GetLocations(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	v := req.URL.Query()

	lat, latErr := strconv.ParseFloat(v.Get("latitude"), 64)
	long, longErr := strconv.ParseFloat(v.Get("longitude"), 64)

	//mev := validators.NewMartaEntitiesValidator()

	if longErr != nil || latErr != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	lt := transformers.NewLocationTransformer()

	response := stationsLocationResponse{lt.GetNearestLocations(lat, long)}

	json.NewEncoder(w).Encode(response)
}
