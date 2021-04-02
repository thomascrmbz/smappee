package smappee

import "time"

type ServiceLocation struct {
	Name               string `json:"name"`
	UUID               string `json:"uuid"`
	ID                 int    `json:"id"`
	DeviceSerialNumber string `json:"device_serial_number"`

	Latitude            float32
	Longtitude          float32
	ElectricityCost     float32
	ElectricityCurrency string
	Timezone            string

	ChannelsConfiguration ChannelsConfiguration

	From time.Time
}

type serviceLocationsResponse struct {
	AppName          string                    `json:"appName"`
	ServiceLocations []serviceLocationResponse `json:"serviceLocations"`
}

type serviceLocationResponse struct {
	Name                  string  `json:"name"`
	UUID                  string  `json:"serviceLocationUuid"`
	ID                    int     `json:"serviceLocationId"`
	DeviceSerialNumber    string  `json:"deviceSerialNumber"`
	Latitude              float32 `json:"lat"`
	Longtitude            float32 `json:"lon"`
	ElectricityCost       float32
	ElectricityCurrency   string
	Timezone              string
	Appliances            []interface{} // TODO
	Actuators             []interface{} // TODO
	Sensors               []interface{} // TODO
	Monitors              []interface{} // TODO
	ChannelsConfiguration ChannelsConfiguration
	Custom                []interface{} // TODO
	From                  int64
}

type ChannelsConfiguration struct {
	InputChannels []InputChannelResponse
}

type InputChannelResponse struct {
	CTInput            int
	Phase              int
	InputChannelType   string
	Reversed           bool
	Nilm               bool
	Balanced           bool
	InputChannelCTType string
}

type ElectricityConsumption struct {
	Timestamp       time.Time
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
	active          []float64
	ReactiveWh      float64
	ReactiveW       float64
	Voltages        [3]float32
	Current         float64
	// LineVoltages, PhaseVoltages

	ctx *context
}

type electricityConsumptionsResponse struct {
	ServiceLocationID int                              `json:"serviceLocationId"`
	Consumptions      []electricityConsumptionResponse `json:"consumptions"`
}

type electricityConsumptionResponse struct {
	Timestamp        int64
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
	ConsumptionW  float64
	ConsumptionWh float64
	Name          string
	Phase         int
}

type ReactiveConsumption struct {
}

type MeteringConfiguration struct {
	Measurements []Measurement
}

type meteringConfigurationResponse struct {
	Measurements []Measurement
	PhaseType    string
}

type Measurement struct {
	ID       int
	Name     string
	Type     string
	Channels []Channel
}

type Channel struct {
	ConsumptionIndex int
	PowerTopicIndex  int
	Name             string
	Phase            string
}
