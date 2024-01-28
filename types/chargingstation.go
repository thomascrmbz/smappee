package types

type ChargingStation struct {
	Id                int
	Name              string
	ServiceLocationId int
}

type ChargingStationConfiguration struct {
	MinimalCurrent int
	MaximalCurrent float64
}
