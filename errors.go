package smappee

import (
	"errors"
)

var (
	ErrorNoDataPoint = errors.New("no datapoint found")
)

func ErrorClientConnection(status string) error {
	return errors.New("Could't connect to smappee: " + status)
}
