package smappee

import (
	"encoding/json"
	"net/url"
	"strconv"
	"time"
)

// GetServiceLocations returns a list of all the service locations to which the specified user account has access to.
func (s *Smappee) GetServiceLocations() *ServiceLocations {
	res, _ := s.newRequest("GET", "/dev/v3/servicelocation", nil)
	serviceLocations := ServiceLocations{}
	json.NewDecoder(res.Body).Decode(&serviceLocations)
	return &serviceLocations
}

type ServiceLocations struct {
	AppName          string             `json:"appName"`
	ServiceLocations *[]ServiceLocation `json:"serviceLocations"`
}

type ServiceLocation struct {
	Name               string `json:"name"`
	UUID               string `json:"serviceLocationUuid"`
	ID                 int    `json:"serviceLocationId"`
	DeviceSerialNumber string `json:"deviceSerialNumber"`
}

// GetElectricityConsumptions returns the electricity consumption on a specific service location during a specified range of time.
// The call support different aggregation levels to obtain different levels of details.
func (s *Smappee) GetElectricityConsumptions(serviceLocationID int, from time.Time, to time.Time, aggregation int) *ElectricityConsumptions {
	parameters := url.Values{}
	parameters.Set("aggregation", strconv.Itoa(aggregation))
	parameters.Set("from", strconv.FormatInt(from.Unix(), 10))
	parameters.Set("to", strconv.FormatInt(to.Unix(), 10))

	res, _ := s.newRequest("GET", "/dev/v3/servicelocation/"+strconv.Itoa(serviceLocationID)+"/consumption", nil, parameters)
	electricityConsumptions := ElectricityConsumptions{}
	json.NewDecoder(res.Body).Decode(&electricityConsumptions)
	return &electricityConsumptions
}

type ElectricityConsumptions struct {
	ServiceLocationID int                       `json:"serviceLocationId"`
	Consumptions      *[]ElectricityConsumption `json:"consumptions"`
}

type ElectricityConsumption struct {
	Timestamp        int
	Consumption      float32
	Solar            float32
	AlwaysOn         float32
	GridImport       float32
	GridExport       float32
	SelfConsumption  float32
	SelfSufficiency  float32
	Active           []float64
	Reactive         []float64
	Voltages         []float32
	LineVoltages     []float32
	LineVoltagesMin  []float32
	LineVoltagesMax  []float32
	PhaseVoltages    []float32
	PhaseVoltagesMin []float32
	PhaseVoltagesMax []float32
	Current          []float32
	CurrentMin       []float32
	CurrentMax       []float32
	CurrentHarmonics [][]float32
	VoltageHarmonics [][]float32
}
