package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/smartatransit/gomarta"
	"github.com/smartatransit/third_rail/pkg/clients"
	"github.com/smartatransit/third_rail/pkg/controllers"
	"github.com/smartatransit/third_rail/pkg/middleware"
	"log"
	"net/http"
)

var martaClient *gomarta.Client
var twitterClient *twitter.Client

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if martaClient == nil {
		martaClient = clients.GetMartaClient()
	}

	if twitterClient == nil {
		twitterClient = clients.GetTwitterClient()
	}

	mountAndServe(martaClient, twitterClient)
}

func mountAndServe(mc *gomarta.Client, tc *twitter.Client) {
	router := mux.NewRouter()

	liveController := controllers.LiveController{MartaClient: mc}
	liveRouter := router.PathPrefix("/live").Subrouter()
	liveRouter.HandleFunc("/schedule/line/{line}", liveController.GetScheduleByLine).Methods("GET")
	liveRouter.HandleFunc("/schedule/station/{station}", liveController.GetScheduleByStation).Methods("GET")

	staticController := controllers.StaticController{}
	staticRouter := router.PathPrefix("/static").Subrouter()
	staticRouter.HandleFunc("/schedule/station", staticController.GetStaticScheduleByStation).Methods("GET")
	staticRouter.HandleFunc("/lines", staticController.GetLines).Methods("GET")
	staticRouter.HandleFunc("/directions", staticController.GetDirections).Methods("GET")
	staticRouter.HandleFunc("/stations", staticController.GetStations).Methods("GET")
	staticRouter.HandleFunc("/stations/location", staticController.GetLocations).Methods("GET")

	smartController := controllers.SmartController{TwitterClient: tc}
	smartRouter := router.PathPrefix("/smart").Subrouter()
	smartRouter.HandleFunc("/parking", smartController.GetParkingStatus).Methods("GET")

	fmt.Println("started on port :5000")

	log.Fatal(http.ListenAndServe(":5000", handlers.CORS(
		handlers.AllowCredentials(),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"POST", "PUT", "DELETE", "PATCH", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With"}),
	)(middleware.SuffixMiddleware(router))))
}
