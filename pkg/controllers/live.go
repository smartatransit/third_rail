package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/smartatransit/third_rail/pkg/clients"
	"github.com/smartatransit/third_rail/pkg/transformers"
	"github.com/smartatransit/third_rail/pkg/validators"
	"log"
	"net/http"
)

type LiveController struct {
	MartaClient clients.MartaClient
}

func (controller LiveController) GetScheduleByLine(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	line := params["line"]

	events, _ := controller.MartaClient.GetTrains()
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

func (controller LiveController) GetScheduleByStation(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	station := params["station"]

	log.Printf("Displaying schedules for %s", station)

	events, _ := controller.MartaClient.GetTrains()
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
