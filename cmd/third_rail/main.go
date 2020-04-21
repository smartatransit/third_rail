package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"
	_ "github.com/smartatransit/third_rail/docs"
	"github.com/smartatransit/third_rail/pkg/clients"
	"github.com/smartatransit/third_rail/pkg/clients/marta_client"
	"github.com/smartatransit/third_rail/pkg/clients/twitter_client"
	"github.com/smartatransit/third_rail/pkg/controllers"
	"github.com/smartatransit/third_rail/pkg/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

var martaClient clients.MartaClient
var twitterClient clients.TwitterClient

func main() {
	var opts options
	_, err := flags.Parse(&opts)

	if err != nil {
		log.Fatal(err)
	}

	if martaClient == nil {
		martaClient = marta_client.GetMartaClient(opts.MartaAPIKey, opts.MartaCacheTTL)
	}

	if twitterClient == nil {
		twitterClient = twitter_client.GetTwitterClient(opts.TwitterClientID, opts.TwitterClientSecret, opts.TwitterCacheTTL)
	}

	mountAndServe(martaClient, twitterClient)
}

// @title SMARTA API
// @version 1.0
// @description API to serve you SMARTA data

// @contact.name SMARTA Support
// @contact.email smartatransit@gmail.com

// @license.name GNU General Public License v3.0
// @license.url https://github.com/smartatransit/third_rail/blob/master/LICENSE

// @host third-rail.services.ataper.net

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func mountAndServe(mc clients.MartaClient, tc clients.TwitterClient) {
	router := mux.NewRouter()

	liveController := controllers.LiveController{MartaClient: mc}
	liveRouter := router.PathPrefix("/live").Subrouter()
	liveRouter.HandleFunc("/schedule/line/{line}", liveController.GetScheduleByLine).Methods("GET")
	liveRouter.HandleFunc("/schedule/station/{station}", liveController.GetScheduleByStation).Methods("GET")
	liveRouter.HandleFunc("/alerts", liveController.GetAlerts).Methods("GET")

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

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	fmt.Println("started on port :5000")

	log.Fatal(http.ListenAndServe(":5000", handlers.CORS(
		handlers.AllowCredentials(),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"POST", "PUT", "DELETE", "PATCH", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With"}),
	)(middleware.SuffixMiddleware(router))))
}
