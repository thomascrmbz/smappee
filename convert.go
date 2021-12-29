package smappee

import "time"

func convertServiceLocation(ctx *context, sli serviceLocationResponse) ServiceLocation {
	return ServiceLocation{
		Name:                sli.Name,
		UUID:                sli.UUID,
		ID:                  sli.ID,
		DeviceSerialNumber:  sli.DeviceSerialNumber,
		Latitude:            sli.Latitude,
		Longtitude:          sli.Longtitude,
		ElectricityCost:     sli.ElectricityCost,
		ElectricityCurrency: sli.ElectricityCurrency,
		Timezone:            sli.Timezone,

		ChannelsConfiguration: sli.ChannelsConfiguration,

		From: time.Unix(0, sli.From*int64(time.Millisecond)),

		ctx: ctx,
	}
}

func convertElectricityConsumption(ctx *context, c electricityConsumptionResponse) ElectricityConsumption {
	return ElectricityConsumption{
		Timestamp:       time.Unix(0, c.Timestamp*int64(time.Millisecond)),
		ConsumptionWh:   c.Consumption,
		ConsumptionW:    c.Consumption * 12,
		SolarWh:         c.Solar,
		SolarW:          c.Solar * 12,
		AlwaysOnWh:      c.AlwaysOn,
		AlwaysOnW:       c.AlwaysOn * 12,
		GridImportWh:    c.GridImport,
		GridImportW:     c.GridImport * 12,
		GridExportWh:    c.GridExport,
		GridExportW:     c.GridExport * 12,
		SelfSufficiency: c.SelfSufficiency,
		SelfConsumption: c.SelfConsumption,
		active:          c.Active,
		ActiveWh:        sum(c.Active),
		ActiveW:         round(sum(c.Active) * 12),
		ReactiveWh:      sum(c.Reactive),
		ReactiveW:       round(sum(c.Reactive) * 12),
		Voltages:        c.Voltages,
		Current:         sum(c.Current),

		ctx: ctx,
	}
}

func convertSensorConsumption(sensor sensorConsumptionResponse) SensorConsumption {
	return SensorConsumption{
		Timestamp: time.Unix(0, sensor.Timestamp*int64(time.Millisecond)),

		Value1: sensor.Value1,
		Value2: sensor.Value2,
		Value3: sensor.Value3,
		Value4: sensor.Value4,

		Temperature: sensor.Temperature,
		Humidity:    sensor.Humidity,
		Battery:     sensor.Battery,
	}
}
