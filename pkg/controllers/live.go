package controllers

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/smartatransit/third_rail/pkg/clients"
	"github.com/smartatransit/third_rail/pkg/transformers"
	"github.com/smartatransit/third_rail/pkg/validators"
)

type LiveController struct {
	MartaClient clients.MartaClient
}

// GetScheduleByLine godoc
// @Summary Get Schedule By Line
// @Description  Given a line, return the current live schedule
// @Param line path string true "RED, GOLD, BLUE, GREEN"
// @Produce  json
// @Success 200 {object} response
// @Router /live/schedule/line/{line} [get]
// @Security ApiKeyAuth
func (controller LiveController) GetScheduleByLine(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	line := params["line"]

	events, _ := controller.MartaClient.GetTrains()
	mev := validators.NewMartaEntitiesValidator()
	eventTransformer := transformers.NewEventTransformer(mev)

	var resp Response

	for _, event := range transformers.FilterByLine(events, line) {
		schedule := eventTransformer.GetSchedule(event)
		station := eventTransformer.GetStation(event)
		resp.Data = append(resp.Data, responseData{Schedule: schedule, Station: station})
	}

	json.NewEncoder(w).Encode(resp)
}

// GetScheduleByStation godoc
// @Summary Get Schedule By Station
// @Description  Given a station, return the current live schedule
// @Param station path string true "TODO: Enter all stations as enum?"
// @Produce  json
// @Success 200 {object} response
// @Router /live/schedule/station/{station} [get]
// @Security ApiKeyAuth
func (controller LiveController) GetScheduleByStation(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	station := params["station"]

	log.Printf("Displaying schedules for %s", station)

	events, _ := controller.MartaClient.GetTrains()
	mev := validators.NewMartaEntitiesValidator()
	eventTransformer := transformers.NewEventTransformer(mev)

	var resp Response

	for _, event := range transformers.FilterByStation(events, station) {
		schedule := eventTransformer.GetSchedule(event)
		station := eventTransformer.GetStation(event)
		resp.Data = append(resp.Data, responseData{Schedule: schedule, Station: station})
	}

	json.NewEncoder(w).Encode(resp)
}

// GetAlerts godoc
// @Summary Get Alerts from various MARTA sources
// @Description MARTA alerts sourced from their official twitter account
// @Produce  json
// @Success 200 {object} alertResponse
// @Router /live/alerts [get]
// @Security ApiKeyAuth
func (controller LiveController) GetAlerts(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	alerts, _ := controller.MartaClient.GetAlerts()

	var resp AlertResponse
	resp.Data = alerts

	json.NewEncoder(w).Encode(resp)
}
