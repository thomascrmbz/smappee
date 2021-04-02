
## Methods

### GetServiceLocations
Returns a list of all the service locations to which the specified user account has access to.

### GetServiceLocation
Returns detailed information of a specific service location.

<!--
### ShareServiceLocation
Shares access to a service location or updates role of a user while accessing the service location. 
-->

### DeleteServiceLocation
Removes a service location.

### UpdateServiceLocation
Modifies a service location.

### CreateServiceLocation
Creates a service location.

### GetElectricityConsumption
Returns the electricity consumption on a specific service location during a specified range of time.

### GetElectricityConsumptions
Returns a list of electricity consumptions on a specific service location during a specified range of time.

### GetSensorConsumption
Returns the consumption of energy on a specific sensor that is active on a specific service location during a specified range of time.

### GetSwitchConsumption

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

### GetMeteringConfiguration
Returns detailed information of a specific service location.

### GetCostAnalysis
Returns cost breakdown of measured energy by the Smappee monitor

### GetEvents
Returns an overview of the appliance events reported on a specic service location. It is possible to filter on one or more appliances.

### GetActuatorState

### SetActuatorState

```go
s.GetServiceLocations()
s.GetServiceLocationInfo()
```