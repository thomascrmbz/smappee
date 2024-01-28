package smappee

import (
	"net/http"

	"thomascrmbz.com/smappee/internal/unlicensed"
	"thomascrmbz.com/smappee/types"
)

type Client interface {
	Authenticate() error
	GetChargingStations() ([]types.ChargingStation, error)
	GetServiceLocations() ([]types.ServiceLocation, error)
	UpdateChargingStationConfiguration(cs types.ChargingStation, configuration types.ChargingStationConfiguration) error
}

func NewUnlicensedClient(username, password string) (Client, error) {
	client := &unlicensed.Client{
		Username:   username,
		Password:   password,
		HttpClient: &http.Client{},
	}
	return client, client.Authenticate()
}
