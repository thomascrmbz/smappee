# Smappee client written in go

This api client does NOT follow the official smappee api documentation.
This is due to the fact that their api is not that user friendly.

[![GoDoc](https://godoc.org/thomascrmbz.com/smappee?status.svg)](https://pkg.go.dev/thomascrmbz.com/smappee/?tab=subdirectories)
[![Go Reportcard](https://goreportcard.com/badge/thomascrmbz.com/smappee)](https://goreportcard.com/report/thomascrmbz.com/smappee)

## Installation
```bash
go get thomascrmbz.com/smappee
```

## Usage
```go
package main

import (
	"fmt"
	"log"
	"time"

	"thomascrmbz.com/smappee"
)

func main() {
	s, err := smappee.NewSmappee(
		"clientID", "clientSecret", "username", "password",
	)

	if err != nil {
		log.Fatal(err)
	}

	c, _ := s.GetElectricityConsumption(60904, time.Now())
	fmt.Println(c)
}
```

## Examples
```go
// Create Smappee client
s, err := smappee.NewSmappee(clientID, clientSecret, username, password)

// Get Service Locations
locations, err := s.GetServiceLocations()

// Get Electricity Consumption
consumption, err := s.GetElectricityConsumption(id)

// Get Active Consumptions
activeConsumptions, err := consumption.GetActiveConsumptions()

// Filter Active Consumptions by channel name
activeConsumptions, err := consumption.GetActiveConsumptions("channel1", "channel2")
```

## Methods
```go
// Service Locations
func (s *Smappee) GetServiceLocations() ([]ServiceLocation, error)
func (s *Smappee) GetServiceLocation(id int) (ServiceLocation, error)
func (s *Smappee) DeleteServiceLocation(id int) (error)

// Electricity Consumption
func (s *Smappee) GetElectricityConsumption(id int, timestamp ...time.Time) (ElectricityConsumption, error)
func (s *Smappee) GetElectricityConsumptions(id int, aggregation int, from time.Time, to ...time.Time) ([]ElectricityConsumption, error)

func (ec *ElectricityConsumption) GetActiveConsumptions(name ...string) ([]ActiveConsumption, error)

// Metering Configuration
func (s *Smappee) GetMeteringConfiguration(id int) (MeteringConfiguration, error)

func (sl *ServiceLocation) GetSensorConsumptions(id int, aggregation int, from time.Time, to ...time.Time) ([]SensorConsumption, error)

```

Read the [reference docs](https://pkg.go.dev/thomascrmbz.com/smappee) for more information about how to use all the methods.

## Contributing
This project is a work in progress.
Any feedback and requests are welcome.
If you feel like something is missing, please open an issue on GitHub.

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Added some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

## License
See the [LICENSE](./LICENSE) file for licensing information.
