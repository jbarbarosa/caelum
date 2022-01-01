package factory

import "caelum/domain/entity"

type PredictionFactory struct {
	newPrediction *entity.Prediction
}

func (p *PredictionFactory) New() *PredictionFactory {
	p.newPrediction = &entity.Prediction{}
	return p
}

func (p *PredictionFactory) AddTemperature(temperature int8) *PredictionFactory {
	p.newPrediction.Temperature = temperature
	return p
}

func (p *PredictionFactory) AddHumidity(humidity int8) *PredictionFactory {
	p.newPrediction.Humidity = humidity
	return p
}

func (p *PredictionFactory) AddPressure(pressure int) *PredictionFactory {
	p.newPrediction.Pressure = pressure
	return p
}

func (p *PredictionFactory) AddThermalSensation(thermalSensation int8) *PredictionFactory {
	p.newPrediction.ThermalSensation = thermalSensation
	return p
}

func (p *PredictionFactory) AddUv(uv int8) *PredictionFactory {
	p.newPrediction.Uv = uv
	return p
}

func (p *PredictionFactory) AddVisibility(visibility int) *PredictionFactory {
	p.newPrediction.Visibility = visibility
	return p
}

func (p *PredictionFactory) AddDate(date int16) *PredictionFactory {
	p.newPrediction.UnixTime = date
	return p
}

func (p *PredictionFactory) AddWind(w entity.Wind) *PredictionFactory {
	p.newPrediction.Wind = w
	return p
}

func (p *PredictionFactory) AddBy(by string) *PredictionFactory {
	p.newPrediction.By = by
	return p
}

func (p *PredictionFactory) Finish() entity.Prediction {
	prediction := *p.newPrediction
	p.newPrediction = nil
	return prediction
}
