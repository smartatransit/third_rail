package controllers

import (
	"encoding/json"
	"github.com/dghubble/go-twitter/twitter"
	"net/http"
)

type SmartController struct {
	TwitterClient *twitter.Client
}

func (controller SmartController) GetParkingStatus(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	v := req.URL.Query()

	stationName := v.Get("station_name")

	if len(stationName) < 1 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	tweets := controller.getParking()

	parkingUpdates := make([]parkingData, len(tweets.Statuses))

	for i, tweet := range tweets.Statuses {
		parkingUpdates[i] = parkingData{Status: tweet.FullText}
	}


	response := parkingResponse{parkingUpdates}

	json.NewEncoder(w).Encode(response)
}

func (controller SmartController) getParking() *twitter.Search {
	search, _, _ := controller.TwitterClient.Search.Tweets(&twitter.SearchTweetParams{
		Query: "#parkingupdate",
		TweetMode: "extended",
	})

	return search

}

