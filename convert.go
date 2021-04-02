package smappee

import "time"

func convertServiceLocation(sli serviceLocationResponse) ServiceLocation {
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
	}
}

func convertElectricityConsumption(c electricityConsumptionResponse) ElectricityConsumption {
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
		ActiveWh:        sum(c.Active),
		ActiveW:         round(sum(c.Active) * 12),
		ReactiveWh:      sum(c.Reactive),
		ReactiveW:       round(sum(c.Reactive) * 12),
		Voltages:        c.Voltages,
		Current:         sum(c.Current),
	}
}
