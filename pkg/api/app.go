package api

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	_ "github.com/smartatransit/third_rail/docs"
	"github.com/smartatransit/third_rail/pkg/clients"
	"github.com/smartatransit/third_rail/pkg/clients/marta_client"
	"github.com/smartatransit/third_rail/pkg/clients/twitter_client"
	"github.com/smartatransit/third_rail/pkg/controllers"
	"github.com/smartatransit/third_rail/pkg/middleware"
	"github.com/smartatransit/third_rail/pkg/models"
	"github.com/smartatransit/third_rail/pkg/seed"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

type App struct {
	Router        *mux.Router
	DB            *gorm.DB
	MartaClient   clients.MartaClient
	TwitterClient clients.TwitterClient
	Options       Options
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
func (app *App) Start(bootstrap bool, customRouter func()) {
	if app.DB == nil {
		dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			app.Options.DbHost,
			app.Options.DbPort,
			app.Options.DbUsername,
			app.Options.DbName,
			app.Options.DbPassword)

		var err error
		app.DB, err = gorm.Open("postgres", dbURI)

		if err != nil {
			log.Fatalf("Could not connect database: %v", err)
		}
	}

	if bootstrap {
		app.DB = models.DBMigrate(app.DB)
		seed.Seed(app.DB)
	}

	//defer app.DB.Close()

	if customRouter != nil {
		customRouter()
	} else {
		if app.MartaClient == nil {
			app.MartaClient = marta_client.GetMartaClient(app.Options.MartaAPIKey, app.Options.MartaCacheTTL)
		}

		if app.TwitterClient == nil {
			app.TwitterClient = twitter_client.GetTwitterClient(app.Options.TwitterClientID, app.Options.TwitterClientSecret, app.Options.TwitterCacheTTL)
		}

		app.Router = mux.NewRouter()
		app.mountLiveRoutes()
		app.MountStaticRoutes()
		app.mountSmartRoutes()
		app.mountRiderRoutes()
		app.mountSwaggerRoutes()

		app.serve()
	}
}

func (app *App) mountLiveRoutes() {
	if app.MartaClient == nil {
		log.Fatal("No MARTA client present - unable to mount Live routes.")
	}

	liveController := controllers.LiveController{MartaClient: app.MartaClient}
	liveRouter := app.Router.PathPrefix("/live").Subrouter()
	liveRouter.HandleFunc("/schedule/line/{line}", liveController.GetScheduleByLine).Methods("GET")
	liveRouter.HandleFunc("/schedule/station/{station}", liveController.GetScheduleByStation).Methods("GET")
	liveRouter.HandleFunc("/alerts", liveController.GetAlerts).Methods("GET")
}

func (app *App) MountStaticRoutes() {
	staticController := controllers.StaticController{}
	staticRouter := app.Router.PathPrefix("/static").Subrouter()
	staticRouter.HandleFunc("/schedule/station", staticController.GetStaticScheduleByStation).Methods("GET")
	staticRouter.HandleFunc("/lines", func(w http.ResponseWriter, r *http.Request) {
		staticController.GetLines(app.DB, w, r)
	}).Methods("GET")
	staticRouter.HandleFunc("/directions", func(w http.ResponseWriter, r *http.Request) {
		staticController.GetDirections(app.DB, w, r)
	}).Methods("GET")
	staticRouter.HandleFunc("/stations", func(w http.ResponseWriter, r *http.Request) {
		staticController.GetStations(app.DB, w, r)
	}).Methods("GET")
	staticRouter.HandleFunc("/stations/location", staticController.GetLocations).Methods("GET")
}

func (app *App) mountSmartRoutes() {
	if app.TwitterClient == nil {
		log.Fatal("No Twitter client present - unable to mount Smart routes.")
	}

	smartController := controllers.SmartController{TwitterClient: app.TwitterClient}
	smartRouter := app.Router.PathPrefix("/smart").Subrouter()
	smartRouter.HandleFunc("/parking", smartController.GetParkingStatus).Methods("GET")
	smartRouter.HandleFunc("/emergencies", smartController.GetEmergencyStatus).Methods("GET")
}

func (app *App) mountRiderRoutes() {
	riderController := controllers.RiderController{}
	riderRouter := app.Router.PathPrefix("/rider").Subrouter()
	riderRouter.HandleFunc("/alerts", func(w http.ResponseWriter, r *http.Request) {
		riderController.GetRiderAlerts(app.DB, w, r)
	}).Methods("GET")
	riderRouter.HandleFunc("/migrate", func(w http.ResponseWriter, r *http.Request) {
		riderController.Migrate(app.DB, w, r)
	}).Methods("GET")
}

func (app *App) mountSwaggerRoutes() {
	app.Router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
}

func (app *App) serve() {
	fmt.Println("started on port :5000")

	log.Fatal(http.ListenAndServe(":5000", handlers.CORS(
		handlers.AllowCredentials(),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"POST", "PUT", "DELETE", "PATCH", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With"}),
	)(middleware.SuffixMiddleware(app.Router))))
}
