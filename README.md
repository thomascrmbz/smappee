# Smappee client written in go

This api client does NOT follow the official smappee api documentation.
This is due to the fact that their api is not that user friendly.

## Example

```golang
// Create Smappee client
s := smappee.NewSmappee(clientID, clientSecret, username, password)

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

### GetServiceLocations
Returns a list of all the service locations to which the specified user account has access to.

### GetServiceLocation
Returns detailed information of a specific service location.

### DeleteServiceLocation
Removes a service location.

<!--
### UpdateServiceLocation
Modifies a service location.

### CreateServiceLocation
Creates a service location.

### ShareServiceLocation
Shares access to a service location or updates role of a user while accessing the service location. 

-->

### GetElectricityConsumption
Returns the electricity consumption on a specific service location during a specified range of time.

### GetElectricityConsumptions
Returns a list of electricity consumptions on a specific service location during a specified range of time.

### GetActiveConsumptions
Returns a list of ActiveConsumptions on a specific ElectricityConsumption

<!--
### GetReactiveConsumptions
Returns a list of ReactiveConsumptions on a specific ElectricityConsumption

### GetSensorConsumption
Returns the consumption of energy on a specific sensor that is active on a specific service location during a specified range of time.

### GetSwitchConsumption
-->

### GetMeteringConfiguration
Returns detailed information of a specific service location.

<!--
### GetSmartDevices
Returns detailed information of a single smart device installed at a specific service location.

### GetSmartDevice
Returns detailed information of the smart devices installed at a specific service location.

### UpdateSmartDevice
Modifies a smart device at a specific service location.

### CreateSmartDevice
Creates a new smart device at a specific service location.

### DeleteSmartDevice
Removes a smart device at a specific service location.
-->

<!--
### GetActuatorState

### SetActuatorState
-->

<!--
### GetCostAnalysis
Returns cost breakdown of measured energy by the Smappee monitor
-->

<!--
### GetEvents
Returns an overview of the appliance events reported on a specic service location. It is possible to filter on one or more appliances.
-->

## Contributing
This project is a work in progress.
Any feedback and requests are welcome.
If you feel like something is missing, please open an issue on GitHub.

## License
See the [LICENSE](./LICENSE) file for licensing information.
