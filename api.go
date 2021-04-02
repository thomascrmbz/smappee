package smappee

import (
	"encoding/json"
	"errors"
	"net/url"
	"strconv"
	"time"
)

func (s *Smappee) GetServiceLocations() ([]ServiceLocation, error) {
	res, err := s.newRequest("GET", "/dev/v3/servicelocation", nil)
	serviceLocationsResponse := serviceLocationsResponse{}
	json.NewDecoder(res.Body).Decode(&serviceLocationsResponse)

	serviceLocations := []ServiceLocation{}

	for _, sl := range serviceLocationsResponse.ServiceLocations {
		serviceLocations = append(serviceLocations, convertServiceLocation(sl))
	}

	return serviceLocations, err
}

func (s *Smappee) GetServiceLocation(id int) (ServiceLocation, error) {
	res, err := s.newRequest("GET", "/dev/v3/servicelocation/"+strconv.Itoa(id)+"/info", nil)
	sli := serviceLocationResponse{}
	json.NewDecoder(res.Body).Decode(&sli)

	return convertServiceLocation(sli), err
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

	res, err := s.newRequest("GET", "/dev/v3/servicelocation/"+strconv.Itoa(id)+"/consumption", nil, parameters)
	electricityConsumptionsResponse := electricityConsumptionsResponse{}
	json.NewDecoder(res.Body).Decode(&electricityConsumptionsResponse)

	if len(electricityConsumptionsResponse.Consumptions) == 0 {
		return ElectricityConsumption{}, errors.New("no datapoint found")
	}

	i := len(electricityConsumptionsResponse.Consumptions) - 1
	c := electricityConsumptionsResponse.Consumptions[i]

	return convertElectricityConsumption(&context{
		Smappee: s,
		ServiceLocation: &ServiceLocation{
			ID: id,
		},
	}, c), err
}

func (s *Smappee) GetElectricityConsumptions(aggregation int, from time.Time, to ...time.Time) ([]ElectricityConsumption, error) {
	electricityConsumptions := []ElectricityConsumption{}
	return electricityConsumptions, nil
}

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

func (ec *ElectricityConsumption) GetReactiveConsumptions() ([]ReactiveConsumption, error) {
	reactiveConsumptions := []ReactiveConsumption{}
	return reactiveConsumptions, nil
}

func (s *Smappee) GetMeteringConfiguration(id int) (MeteringConfiguration, error) {
	res, err := s.newRequest("GET", "/dev/v3/servicelocation/"+strconv.Itoa(id)+"/meteringconfiguration", nil)
	meteringConfigurationResponse := meteringConfigurationResponse{}
	json.NewDecoder(res.Body).Decode(&meteringConfigurationResponse)

	meteringConfiguration := MeteringConfiguration{
		Measurements: meteringConfigurationResponse.Measurements,
	}
	return meteringConfiguration, err
}
