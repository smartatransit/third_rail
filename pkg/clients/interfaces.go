package clients

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/smartatransit/gomarta"
	"github.com/smartatransit/third_rail/pkg/schemas/marta_schemas"
)

type MartaClient interface {
	GetTrains() (gomarta.TrainAPIResponse, error)
	GetAlerts() (marta_schemas.Alerts, error)
}

type TwitterClient interface {
	Search(string, *twitter.SearchTweetParams) (*twitter.Search, error)
}

