package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jessevdk/go-flags"
	"github.com/smartatransit/third_rail/pkg/clients"
	"github.com/smartatransit/third_rail/pkg/controllers"
	"github.com/smartatransit/third_rail/pkg/middleware"
)

var martaClient clients.MartaClient
var twitterClient clients.TwitterClient

type options struct {
	TwitterClientID     string `long:"twitter-client-id" env:"TWITTER_CLIENT_ID" description:"the client id for the twitter acount"`
	TwitterClientSecret string `long:"twitter-client-secret" env:"TWITTER_CLIENT_SECRET" description:"the client secret for the twitter acount"`
	MartaAPIKey         string `long:"marta-api-key" env:"MARTA_API_KEY" description:"marta api key"`
	TwitterCacheTTL     int    `long:"twitter-cache-ttl" env:"TWITTER_CACHE_TTL" description:"how long we keep the twitter responses" required:"true" default:15`
	MartaCacheTTL       int    `long:"marta-cache-ttl" env:"MARTA_CACHE_TTL" description:"how long we keep the marta responses" required:"true" default:15`
}

func main() {
	var opts options
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	if martaClient == nil {
		martaClient = clients.GetMartaClient(opts.MartaAPIKey, opts.MartaCacheTTL)
	}

	if twitterClient == nil {
		twitterClient = clients.GetTwitterClient(opts.TwitterClientID, opts.TwitterClientSecret, opts.TwitterCacheTTL)
	}

	mountAndServe(martaClient, twitterClient)
}

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

	fmt.Println("started on port :5000")

	log.Fatal(http.ListenAndServe(":5000", handlers.CORS(
		handlers.AllowCredentials(),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"POST", "PUT", "DELETE", "PATCH", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With"}),
	)(middleware.SuffixMiddleware(router))))
}
