package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/smartatransit/scrapedumper/pkg/postgres"
	"github.com/smartatransit/third_rail/pkg/clients"
	"github.com/smartatransit/third_rail/pkg/models"
)

type MartaMirror interface {
	GetLatestEstimates(stationID uint) (res []postgres.LastestEstimate, err error)
}

type SmartController struct {
	MartaClient   clients.MartaClient
	TwitterClient clients.TwitterClient

	SDClient MartaMirror
}

// GetScheduleByStation godoc
// @Summary Get Schedule By Station
// @Description  Given a station ID, return the latest estimates for trains
// @Param id path string true "Unique id of the station"
// @Produce json
// @Success 200 {object} Response
// @Router /smart/station/{id} [get]
// @Security ApiKeyAuth
func (controller SmartController) GetStationDetails(db *gorm.DB, w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	v := mux.Vars(req)
	stationId, idErr := strconv.Atoi(v["id"])
	if idErr != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	var station models.Station
	if err := db.First(&station, stationId).Error; err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	sdEstimates, err := controller.SDClient.GetLatestEstimates(uint(stationId))
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var schedules []ScheduleDetail
	for _, estimate := range sdEstimates {
		schedules = append(schedules, convertSDEstimate(estimate))
	}

	response := StationDetailsResponse{
		Data: StationDetails{
			Station:  station,
			Schedule: schedules,
		},
	}

	json.NewEncoder(w).Encode(response)
}

func convertSDEstimate(estimate postgres.LastestEstimate) ScheduleDetail {
	return ScheduleDetail{
		Event: models.ScheduleEvent{
			Destination: models.Station{
				Name: estimate.Destination,
			},
			NextStation: models.Station{
				ID:   *estimate.StationID,
				Name: estimate.Station,
			},
			Direction: models.Direction{
				ID:   *estimate.DirectionID,
				Name: estimate.Direction,
			},
			NextArrival: time.Time(estimate.NextArrival),
		},
	}
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
