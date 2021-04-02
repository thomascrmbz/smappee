package smappee

import (
	"encoding/json"
	"errors"
	"net/url"
	"strconv"
	"time"
)

func (s *Smappee) GetServiceLocations() ([]ServiceLocation, error) {
	res, _ := s.newRequest("GET", "/dev/v3/servicelocation", nil)
	serviceLocationsResponse := serviceLocationsResponse{}
	json.NewDecoder(res.Body).Decode(&serviceLocationsResponse)

	serviceLocations := []ServiceLocation{}

	for _, sl := range serviceLocationsResponse.ServiceLocations {
		serviceLocations = append(serviceLocations, ServiceLocation{
			Name:               sl.Name,
			UUID:               sl.UUID,
			ID:                 sl.ID,
			DeviceSerialNumber: sl.DeviceSerialNumber,
		})
	}

	return serviceLocations, nil
}

func (s *Smappee) GetServiceLocation(id int) (ServiceLocation, error) {
	res, _ := s.newRequest("GET", "/dev/v3/servicelocation/"+strconv.Itoa(id)+"/info", nil)
	sli := serviceLocationInfoResponse{}
	json.NewDecoder(res.Body).Decode(&sli)

	return ServiceLocation{
		Name:                sli.Name,
		UUID:                sli.UUID,
		ID:                  sli.ID,
		DeviceSerialNumber:  sli.DeviceSerialNumber,
		Latitude:            sli.Latitude,
		Longtitude:          sli.Longtitude,
		ElectricityCost:     sli.ElectricityCost,
		ElectricityCurrency: sli.ElectricityCurrency,
		Timezone:            sli.Timezone,

		ChannelsConfiguration: sli.ChannelsConfiguration,

		From: time.Unix(0, sli.From*int64(time.Millisecond)),
	}, nil
}

func (s *Smappee) DeleteServiceLocation(id int) error {
	return nil
}

func (s *Smappee) UpdateServiceLocation(id int, sl ServiceLocation) (ServiceLocation, error) {
	serviceLocation := ServiceLocation{}
	return serviceLocation, nil
}

func (s *Smappee) CreateServiceLocation(sl ServiceLocation) (ServiceLocation, error) {
	serviceLocation := ServiceLocation{}
	return serviceLocation, nil
}

func (s *Smappee) GetElectricityConsumption(id int, timestamp ...time.Time) (ElectricityConsumption, error) {

	to := time.Now()
	if len(timestamp) > 0 {
		to = timestamp[0]
	}
	from := to.Add(-15 * time.Minute)

	parameters := url.Values{}
	parameters.Set("aggregation", strconv.Itoa(1))
	parameters.Set("from", strconv.FormatInt(from.Unix(), 10))
	parameters.Set("to", strconv.FormatInt(to.Unix(), 10))

	res, _ := s.newRequest("GET", "/dev/v3/servicelocation/"+strconv.Itoa(id)+"/consumption", nil, parameters)
	electricityConsumptionsResponse := electricityConsumptionsResponse{}
	json.NewDecoder(res.Body).Decode(&electricityConsumptionsResponse)

	if len(electricityConsumptionsResponse.Consumptions) == 0 {
		return ElectricityConsumption{}, errors.New("no datapoint found")
	}

	i := len(electricityConsumptionsResponse.Consumptions) - 1
	c := electricityConsumptionsResponse.Consumptions[i]

	return ElectricityConsumption{
		Timestamp:       time.Unix(0, c.Timestamp*int64(time.Millisecond)),
		ConsumptionWh:   c.Consumption,
		ConsumptionW:    c.Consumption * 12,
		SolarWh:         c.Solar,
		SolarW:          c.Solar * 12,
		AlwaysOnWh:      c.AlwaysOn,
		AlwaysOnW:       c.AlwaysOn * 12,
		GridImportWh:    c.GridImport,
		GridImportW:     c.GridImport * 12,
		GridExportWh:    c.GridExport,
		GridExportW:     c.GridExport * 12,
		SelfSufficiency: c.SelfSufficiency,
		SelfConsumption: c.SelfConsumption,
		ActiveWh:        sum(c.Active),
		ActiveW:         round(sum(c.Active) * 12),
		ReactiveWh:      sum(c.Reactive),
		ReactiveW:       round(sum(c.Reactive) * 12),
		Voltages:        c.Voltages,
		Current:         sum(c.Current),
	}, nil
}

func (s *Smappee) GetElectricityConsumptions(aggregation int, from time.Time, to ...time.Time) ([]ElectricityConsumption, error) {
	electricityConsumptions := []ElectricityConsumption{}
	return electricityConsumptions, nil
}

func (ec *ElectricityConsumption) GetActiveConsumption() ([]ActiveConsumption, error) {
	activeConsumptions := []ActiveConsumption{}
	return activeConsumptions, nil
}

func (ec *ElectricityConsumption) GetReactiveConsumption() ([]ReactiveConsumption, error) {
	reactiveConsumptions := []ReactiveConsumption{}
	return reactiveConsumptions, nil
}

func (s *Smappee) GetMeteringConfiguration(id int) (MeteringConfiguration, error) {
	meteringConfiguration := MeteringConfiguration{}
	return meteringConfiguration, nil
}
