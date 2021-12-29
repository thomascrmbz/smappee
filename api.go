package smappee

import (
	"encoding/json"
	"net/url"
	"strconv"
	"time"
)

// GetServiceLocations returns a list of all the service locations to which the specified user account has access to.
func (s *Smappee) GetServiceLocations() ([]ServiceLocation, error) {
	res, err := s.newRequest("GET", "/dev/v3/servicelocation", nil)
	serviceLocationsResponse := serviceLocationsResponse{}
	json.NewDecoder(res.Body).Decode(&serviceLocationsResponse)

	serviceLocations := []ServiceLocation{}

	for _, sl := range serviceLocationsResponse.ServiceLocations {
		serviceLocations = append(serviceLocations, convertServiceLocation(&context{Smappee: s}, sl))
	}

	return serviceLocations, err
}

// GetServiceLocation returns detailed information of a specific service location
func (s *Smappee) GetServiceLocation(id int) (ServiceLocation, error) {
	res, err := s.newRequest("GET", "/dev/v3/servicelocation/"+strconv.Itoa(id)+"/info", nil)
	sli := serviceLocationResponse{}
	json.NewDecoder(res.Body).Decode(&sli)

	return convertServiceLocation(&context{Smappee: s}, sli), err
}

// DeleteServiceLocation deletes a service location.
func (s *Smappee) DeleteServiceLocation(id int) error {
	_, err := s.newRequest("DELETE", "/dev/v2/servicelocation/"+strconv.Itoa(id), nil)
	return err
}

// UpdateServiceLocation updates the attributes of a service location that are specified.
func (s *Smappee) UpdateServiceLocation(id int, sl ServiceLocation) (ServiceLocation, error) {
	serviceLocation := ServiceLocation{}
	return serviceLocation, nil
}

// CreateServiceLocation creates a service location.
func (s *Smappee) CreateServiceLocation(sl ServiceLocation) (ServiceLocation, error) {
	serviceLocation := ServiceLocation{}
	return serviceLocation, nil
}

// GetElectricityConsumption returns the electricity consumption on a specific service location during a specified range of time.
func (s *Smappee) GetElectricityConsumption(id int, timestamp ...time.Time) (ElectricityConsumption, error) {

	to := time.Now()
	if len(timestamp) > 0 {
		to = timestamp[0]
	}
	from := to.Add(-15 * time.Minute)

	consumptions, err := s.GetElectricityConsumptions(id, 1, from)

	return consumptions[len(consumptions)-1], err
}

// GetElectricityConsumptions returns a list of electricity consumptions on a specific service location during a specified range of time.
func (s *Smappee) GetElectricityConsumptions(id int, aggregation int, from time.Time, to ...time.Time) ([]ElectricityConsumption, error) {
	toTime := time.Now()
	if len(to) > 0 {
		toTime = to[0]
	}

	parameters := url.Values{}
	parameters.Set("aggregation", strconv.Itoa(aggregation))
	parameters.Set("from", strconv.FormatInt(from.UnixNano()/1e6, 10))
	parameters.Set("to", strconv.FormatInt(toTime.UnixNano()/1e6, 10))

	res, err := s.newRequest("GET", "/dev/v3/servicelocation/"+strconv.Itoa(id)+"/consumption", nil, parameters)
	electricityConsumptionsResponse := electricityConsumptionsResponse{}
	json.NewDecoder(res.Body).Decode(&electricityConsumptionsResponse)

	if len(electricityConsumptionsResponse.Consumptions) == 0 {
		return []ElectricityConsumption{}, ErrorNoDataPoint
	}

	electricityConsumptions := []ElectricityConsumption{}

	for _, c := range electricityConsumptionsResponse.Consumptions {
		electricityConsumptions = append(electricityConsumptions, convertElectricityConsumption(&context{
			Smappee: s,
			ServiceLocation: &ServiceLocation{
				ID: id,
			},
		}, c))
	}

	return electricityConsumptions, err
}

// GetActiveConsumptions returns a list of ActiveConsumption on a specific ElectricityConsumption
func (ec *ElectricityConsumption) GetActiveConsumptions(name ...string) ([]ActiveConsumption, error) {
	mc, err := ec.ctx.Smappee.GetMeteringConfiguration(ec.ctx.ServiceLocation.ID)

	activeConsumptions := []ActiveConsumption{}

	for _, m := range mc.Measurements {
		for _, channel := range m.Channels {
			if !(len(name) > 0 && !stringInSlice(channel.Name, name)) {
				phase := 0

				switch channel.Phase {
				case "PHASEA":
					phase = 0
				case "PHASEB":
					phase = 1
				case "PHASEC":
					phase = 2
				}

				activeConsumptions = append(activeConsumptions, ActiveConsumption{
					ConsumptionW:  round(ec.active[channel.ConsumptionIndex] * 12),
					ConsumptionWh: ec.active[channel.ConsumptionIndex],
					Name:          channel.Name,
					Phase:         phase,
				})
			}
		}
	}

	return activeConsumptions, err
}

// GetReactiveConsumptions returns a list of ReactiveConsumption on a specific ElectricityConsumption
func (ec *ElectricityConsumption) GetReactiveConsumptions() ([]ReactiveConsumption, error) {
	reactiveConsumptions := []ReactiveConsumption{}
	return reactiveConsumptions, nil
}

// GetMeteringConfiguration returns detailed information of a specific service location.
func (s *Smappee) GetMeteringConfiguration(id int) (MeteringConfiguration, error) {
	res, err := s.newRequest("GET", "/dev/v3/servicelocation/"+strconv.Itoa(id)+"/meteringconfiguration", nil)
	meteringConfigurationResponse := meteringConfigurationResponse{}
	json.NewDecoder(res.Body).Decode(&meteringConfigurationResponse)

	meteringConfiguration := MeteringConfiguration{
		Measurements: meteringConfigurationResponse.Measurements,
	}
	return meteringConfiguration, err
}

// GetSensorConsumptions returns the consumption of energy on a specific sensor that is active on a specific service location during a specified range of time.
func (sl *ServiceLocation) GetSensorConsumptions(id int, aggregation int, from time.Time, to ...time.Time) ([]SensorConsumption, error) {
	toTime := time.Now()
	if len(to) > 0 {
		toTime = to[0]
	}

	parameters := url.Values{}
	parameters.Set("aggregation", strconv.Itoa(aggregation))
	parameters.Set("from", strconv.FormatInt(from.UnixNano()/1e6, 10))
	parameters.Set("to", strconv.FormatInt(toTime.UnixNano()/1e6, 10))
	parameters.Set("fillGaps", "false")

	res, err := sl.ctx.Smappee.newRequest("GET", "/dev/v3/servicelocation/"+strconv.Itoa(sl.ID)+"/sensor/"+strconv.Itoa(id)+"/consumption", nil, parameters)
	sensorConsumptionsResponse := sensorConsumptionsResponse{}
	json.NewDecoder(res.Body).Decode(&sensorConsumptionsResponse)

	if len(sensorConsumptionsResponse.Records) == 0 {
		return []SensorConsumption{}, ErrorNoDataPoint
	}

	sensorConsumptions := []SensorConsumption{}

	for _, scr := range sensorConsumptionsResponse.Records {
		sensorConsumptions = append(sensorConsumptions, convertSensorConsumption(scr))
	}

	return sensorConsumptions, err
}
