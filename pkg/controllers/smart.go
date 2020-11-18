package controllers

import (
	"encoding/json"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"github.com/smartatransit/third_rail/pkg/clients"
	"github.com/smartatransit/third_rail/pkg/models"
	"net/http"
	"strconv"
)

type SmartController struct {
	MartaClient   clients.MartaClient
	TwitterClient clients.TwitterClient
}

func (controller SmartController) GetStationDetails(db *gorm.DB, w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	v := mux.Vars(req)

	stationId, idErr := strconv.Atoi(v["id"])
	if idErr != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	var station models.Station
	db.Preload("Detail").Preload("Lines").Preload("Aliases").First(&station, stationId)

	scheduleEvents, realTimeDetails, seErr := models.GetScheduleEventsByStationRealTime(stationId, db)

	if seErr != nil {
		log.Error(seErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var schedules []ScheduleDetail

	for index, event := range scheduleEvents {
		schedules = append(schedules, ScheduleDetail{
			Event:    event,
			RealTime: &realTimeDetails[index],
			Static:   nil,
		})
	}

	response := StationDetailsResponse{
		Data: StationDetails{
			Station:  station,
			Schedule: schedules,
		},
	}

	json.NewEncoder(w).Encode(response)
}

// GetParkingStatus godoc
// @Summary Get Parking Information
// @Description  Get available parking information as informed by twitter
// @Produce  json
// @Success 200 {object} ParkingResponse
// @Router /smart/parking [get]
// @Security ApiKeyAuth
func (controller SmartController) GetParkingStatus(db *gorm.DB, w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	tweets := controller.getParking()

	parkingUpdates := make([]parkingData, len(tweets.Statuses))

	for i, tweet := range tweets.Statuses {
		//TODO : Fuzzy match topic to bind tweet update to specific entity (station/line/direction)
		parkingUpdates[i] = parkingData{Status: tweet.FullText}
	}

	response := ParkingResponse{parkingUpdates}

	json.NewEncoder(w).Encode(response)
}

func (controller SmartController) GetEmergencyStatus(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	tweets := controller.getEmergencies()

	parkingUpdates := make([]parkingData, len(tweets.Statuses))

	for i, tweet := range tweets.Statuses {
		parkingUpdates[i] = parkingData{Status: tweet.FullText}
	}

	response := ParkingResponse{parkingUpdates}

	json.NewEncoder(w).Encode(response)
}

func (controller SmartController) getParking() *twitter.Search {
	search, _ := controller.TwitterClient.Search("parking", &twitter.SearchTweetParams{
		Query:     "from:@martaservice #parkingupdate",
		TweetMode: "extended",
	})

	return search
}

func (controller SmartController) getEmergencies() *twitter.Search {
	search, _ := controller.TwitterClient.Search("emergencies", &twitter.SearchTweetParams{
		Query:     "from:@martapolice",
		TweetMode: "extended",
	})

	return search
}
