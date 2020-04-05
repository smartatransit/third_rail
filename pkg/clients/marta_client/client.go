package marta_client

import (
	"encoding/xml"
	"github.com/karlseguin/ccache"
	log "github.com/sirupsen/logrus"
	"github.com/smartatransit/gomarta"
	"github.com/smartatransit/third_rail/pkg/schemas/marta_schemas"
	"io/ioutil"
	"net/http"
	"time"
)

const MARTA_ALERT_ENDPOINT = "https://martaalerts.com/webdata.aspx"

type MartaAPIClient struct {
	client   *gomarta.Client
	cache    *ccache.Cache
	cacheTTL time.Duration
}

func GetMartaClient(apiKey string, cacheTTL int) MartaAPIClient {
	if len(apiKey) < 1 {
		log.Fatal("No MARTA API key found - real-time API results are unavailable.")
	}

	if cacheTTL < 1 {
		log.Warn("MARTA API caching set to < 1ms - API results are not being cached.")
	}

	var cache = ccache.New(ccache.Configure().MaxSize(1000).ItemsToPrune(100))
	var marta = gomarta.NewDefaultClient(apiKey)

	return MartaAPIClient{client: marta, cache: cache, cacheTTL: time.Duration(cacheTTL)}
}

func (m MartaAPIClient) GetTrains() (gomarta.TrainAPIResponse, error) {
	log.Print("Fetching trains (no cache)")
	trains, err := m.cache.Fetch("trains", time.Second*m.cacheTTL, func() (interface{}, error) {
		return m.client.GetTrains()
	})

	if err != nil {
		return nil, err
	}

	return trains.Value().(gomarta.TrainAPIResponse), nil
}

func (m MartaAPIClient) GetAlerts() (marta_schemas.Alerts, error) {

	alerts, err := m.cache.Fetch("trains", time.Second*m.cacheTTL, func() (interface{}, error) {
		log.Print("Fetching alerts (no cache)")
		resp, err := http.Get(MARTA_ALERT_ENDPOINT)
		if err != nil {
			log.Fatal("Error getting response. ", err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal("Error reading response. ", err)
		}

		var alerts marta_schemas.Alerts
		err = xml.Unmarshal(body, &alerts)

		return alerts, err
	})

	return alerts.Value().(marta_schemas.Alerts), err
}
