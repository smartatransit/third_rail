package marta_client

import (
	"github.com/smartatransit/gomarta"
	"github.com/smartatransit/third_rail/pkg/schemas/marta_schemas"
)

func CreateMartaAPITestClient(h MartaAPIHarness) MartaAPITestClient {
	return MartaAPITestClient{h}
}

type MartaAPITestClient struct {
	harness MartaAPIHarness
}

func (mac MartaAPITestClient) GetTrains() (gomarta.TrainAPIResponse, error) {
	return mac.harness.GetTrains()
}

func (mac MartaAPITestClient) GetAlerts() (marta_schemas.Alerts, error) {
	return mac.harness.GetAlerts()
}

type MartaAPIHarness struct {
	GetTrains func() (gomarta.TrainAPIResponse, error)
	GetAlerts func() (marta_schemas.Alerts, error)
}
