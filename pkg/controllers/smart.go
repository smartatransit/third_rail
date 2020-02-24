package controllers

import (
	"encoding/json"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/smartatransit/third_rail/pkg/clients"
	"net/http"
)

type SmartController struct {
	TwitterClient clients.TwitterClient
}

func (controller SmartController) GetParkingStatus(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//v := req.URL.Query()

	//stationName := v.Get("station_name")

	//if len(stationName) < 1 {
	//	w.WriteHeader(http.StatusUnprocessableEntity)
	//	return
	//}

	tweets := controller.getParking()

	parkingUpdates := make([]parkingData, len(tweets.Statuses))

	for i, tweet := range tweets.Statuses {
		parkingUpdates[i] = parkingData{Status: tweet.FullText}
	}

	response := parkingResponse{parkingUpdates}

	json.NewEncoder(w).Encode(response)
}

func (controller SmartController) GetEmergencyStatus(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	tweets := controller.getEmergencies()

	parkingUpdates := make([]parkingData, len(tweets.Statuses))

	for i, tweet := range tweets.Statuses {
		parkingUpdates[i] = parkingData{Status: tweet.FullText}
	}

	response := parkingResponse{parkingUpdates}

	json.NewEncoder(w).Encode(response)
}


func (controller SmartController) getParking() *twitter.Search {
	search, _ := controller.TwitterClient.Search("parking", &twitter.SearchTweetParams{
		Query: "from:@martaservice #parkingupdate",
		TweetMode: "extended",
	})

	return search
}

func (controller SmartController) getEmergencies() *twitter.Search {
	search, _ := controller.TwitterClient.Search("emergencies", &twitter.SearchTweetParams{
		Query: "from:@martapolice",
		TweetMode: "extended",
	})

	return search
}

