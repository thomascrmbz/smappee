package smappee

type ServiceLocation struct {
	Name               string `json:"name"`
	UUID               string `json:"serviceLocationUuid"`
	ID                 int    `json:"serviceLocationId"`
	DeviceSerialNumber string `json:"deviceSerialNumber"`
}

type serviceLocationsResponse struct {
	AppName          string            `json:"appName"`
	ServiceLocations []ServiceLocation `json:"serviceLocations"`
}

type serviceLocationResponse struct {
	Name               string `json:"name"`
	UUID               string `json:"serviceLocationUuid"`
	ID                 int    `json:"serviceLocationId"`
	DeviceSerialNumber string `json:"deviceSerialNumber"`
}

type ElectricityConsumption struct {
	Timestamp       int
	ConsumptionWh   float32
	ConsumptionW    float32
	SolarWh         float32
	SolarW          float32
	AlwaysOnWh      float32
	AlwaysOnW       float32
	GridImportWh    float32
	GridImportW     float32
	GridExportWh    float32
	GridExportW     float32
	SelfSufficiency float32
	SelfConsumption float32
	ActiveWh        float64
	ActiveW         float64
	ReactiveWh      float64
	ReactiveW       float64
	Voltages        [3]float32
	Current         float64
	// LineVoltages, PhaseVoltages
}

type electricityConsumptionsResponse struct {
	ServiceLocationID int                              `json:"serviceLocationId"`
	Consumptions      []electricityConsumptionResponse `json:"consumptions"`
}

type electricityConsumptionResponse struct {
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
	Voltages         [3]float32
	LineVoltages     []float32
	LineVoltagesMin  []float32
	LineVoltagesMax  []float32
	PhaseVoltages    []float32
	PhaseVoltagesMin []float32
	PhaseVoltagesMax []float32
	Current          []float64
	CurrentMin       []float32
	CurrentMax       []float32
	CurrentHarmonics [][]float32
	VoltageHarmonics [][]float32
}

type ActiveConsumption struct {
}

type ReactiveConsumption struct {
}

type MeteringConfiguration struct {
}
