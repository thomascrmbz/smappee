package unlicensed

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"

	"thomascrmbz.com/smappee/types"
)

type chargingStationsResponse []struct {
	Available       bool      `json:"available"`
	Chargers        []charger `json:"chargers"`
	LastSeen        int       `json:"lastInterval"`
	Parent          location  `json:"parent"`
	ServiceLocation location  `json:"serviceLocation"`
}

type charger struct {
	Id              int    `json:"id"`
	ChargingMode    string `json:"chargingMode"`
	ConnectedStatus string `json:"connectedStatus"`
	Position        int    `json:"position"`
	SerialNumber    string `json:"serialNumber"`
}

type location struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (c *Client) GetChargingStations() ([]types.ChargingStation, error) {
	res, err := c.newRequest(http.MethodGet, "https://dashboard.smappee.net/api/v10/user/chargingstations", nil)

	var response chargingStationsResponse
	json.NewDecoder(res.Body).Decode(&response)

	chargingStations := []types.ChargingStation{}
	for _, cs := range response {
		for _, c := range cs.Chargers {
			chargingStations = append(chargingStations, types.ChargingStation{
				Id:                c.Id,
				Name:              cs.ServiceLocation.Name,
				ServiceLocationId: cs.ServiceLocation.Id,
			})
		}
	}

	return chargingStations, err
}

type updateChargingStationConfigurationReq struct {
	Properties []any `json:"configurationProperties"`
}

type maxCurrentProperty struct {
	Values []maxCurrentPropertyValue `json:"values"`
	Spec   propertySpec              `json:"spec"`
}

type propertySpec struct {
	Name string `json:"name"`
}

type maxCurrentPropertyValue struct {
	Quantity quantityValue `json:"Quantity"`
}

type quantityValue struct {
	Value int    `json:"value"`
	Unit  string `json:"unit"`
}

type minCurrentProperty struct {
}

type percentageProperty struct {
	Values []percentagePropertyValue `json:"values"`
	Spec   propertySpec              `json:"spec"`
}
type percentagePropertyValue struct {
	Integer int `json:"Integer"`
}

func (c *Client) UpdateChargingStationConfiguration(cs types.ChargingStation, configuration types.ChargingStationConfiguration) error {
	req := updateChargingStationConfigurationReq{Properties: []any{}}

	if configuration.MinimalCurrent > 0 {
		// @todo implement
		// req.Properties = append(req.Properties, minCurrentProperty{})
	}
	if configuration.MaximalCurrent > 0 {
		mc := int(math.Ceil(configuration.MaximalCurrent))
		percentage := int((configuration.MaximalCurrent / float64(mc)) * 100)

		req.Properties = append(req.Properties, maxCurrentProperty{
			Spec: propertySpec{
				Name: "etc.smart.device.type.car.charger.config.max.current",
			},
			Values: []maxCurrentPropertyValue{{Quantity: quantityValue{
				Value: mc,
				Unit:  "A",
			}}},
		}, percentageProperty{
			Spec: propertySpec{
				Name: "etc.smart.device.type.car.charger.config.min.excesspct",
			},
			Values: []percentagePropertyValue{{Integer: percentage}},
		})
	}

	res, err := c.newRequest(
		http.MethodPatch,
		fmt.Sprintf(
			"https://dashboard.smappee.net/api/v10/servicelocation/%d/homecontrol/smart/devices/CARCHARGER-%d",
			cs.ServiceLocationId,
			cs.Id,
		),
		req,
	)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return errors.New(res.Status)
	}

	return err
}
