package smappee

import (
	"encoding/json"
	"net/url"
	"strconv"
	"time"
)

func (s *Smappee) GetServiceLocations() ([]ServiceLocation, error) {
	// 	res, _ := s.newRequest("GET", "/dev/v3/servicelocation", nil)
	// 	serviceLocations := ServiceLocations{}
	// 	json.NewDecoder(res.Body).Decode(&serviceLocations)
	// 	return serviceLocations
	serviceLocations := []ServiceLocation{}
	return serviceLocations, nil
}

func (s *Smappee) GetServiceLocation(id int) (ServiceLocation, error) {
	serviceLocation := ServiceLocation{}
	return serviceLocation, nil
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
	from := time.Now().Add(-15 * time.Minute)
	to := time.Now()

	if len(timestamp) > 0 {
		from = timestamp[0]
		to = from.Add(15 * time.Minute)
	}

	parameters := url.Values{}
	parameters.Set("aggregation", strconv.Itoa(1))
	parameters.Set("from", strconv.FormatInt(from.Unix(), 10))
	parameters.Set("to", strconv.FormatInt(to.Unix(), 10))

	res, _ := s.newRequest("GET", "/dev/v3/servicelocation/"+strconv.Itoa(id)+"/consumption", nil, parameters)
	electricityConsumptionsResponse := electricityConsumptionsResponse{}
	json.NewDecoder(res.Body).Decode(&electricityConsumptionsResponse)

	i := len(electricityConsumptionsResponse.Consumptions) - 1
	c := electricityConsumptionsResponse.Consumptions[i]

	electricityConsumption := ElectricityConsumption{
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
	}

	return electricityConsumption, nil
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
