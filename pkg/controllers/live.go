package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/smartatransit/gomarta"
	"github.com/smartatransit/third_rail/pkg/transformers"
	"github.com/smartatransit/third_rail/pkg/validators"
	"log"
	"net/http"
	"os"
)

func GetScheduleByLine(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	line := params["line"]

	martaClient := gomarta.NewDefaultClient(os.Getenv("MARTA_API_KEY"))

	events, _ := martaClient.GetTrains()
	mev := validators.NewMartaEntitiesValidator()
	eventTransformer := transformers.NewEventTransformer(mev)

	var resp response

	for _, event := range transformers.FilterByLine(events, line) {
		schedule := eventTransformer.GetSchedule(event)
		station := eventTransformer.GetStation(event)
		resp.Data = append(resp.Data, responseData{Schedule: schedule, Station: station})
	}

	json.NewEncoder(w).Encode(resp)
}

func GetScheduleByStation(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	station := params["station"]

	log.Printf("Displaying schedules for %s", station)

	martaClient := gomarta.NewDefaultClient(os.Getenv("MARTA_API_KEY"))

	events, _ := martaClient.GetTrains()
	mev := validators.NewMartaEntitiesValidator()
	eventTransformer := transformers.NewEventTransformer(mev)

	var resp response

	for _, event := range transformers.FilterByStation(events, station) {
		schedule := eventTransformer.GetSchedule(event)
		station := eventTransformer.GetStation(event)
		resp.Data = append(resp.Data, responseData{Schedule: schedule, Station: station})
	}

	json.NewEncoder(w).Encode(resp)
}
