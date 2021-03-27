package smappee

import (
	"encoding/json"
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
