package smappee_test

import (
	"os"
	"testing"

	"thomascrmbz.com/smappee"
	"thomascrmbz.com/smappee/types"
)

func TestUnlicensedClient(t *testing.T) {
	if _, err := smappee.NewUnlicensedClient(os.Getenv("SMAPPEE_USERNAME"), os.Getenv("SMAPPEE_PASSWORD")); err != nil {
		t.Error(err)
	}
}

func TestGetAllChargingStations(t *testing.T) {
	client, _ := smappee.NewUnlicensedClient(os.Getenv("SMAPPEE_USERNAME"), os.Getenv("SMAPPEE_PASSWORD"))
	stations, _ := client.GetChargingStations()

	station := stations[0]

	if err := client.UpdateChargingStationConfiguration(station, types.ChargingStationConfiguration{
		MinimalCurrent: 6,
		MaximalCurrent: 32,
	}); err != nil {
		t.Error(err)
	}
}
