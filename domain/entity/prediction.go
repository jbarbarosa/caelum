package entity

type Prediction struct {
	UnixTime         int16
	Temperature      int8
	Uv               int8
	Pressure         int
	Visibility       int
	Humidity         int8
	ThermalSensation int8
	Wind             Wind
	Conditions       Conditions
	By               string
}

func (p *Prediction) ForcePositive() *Prediction {
	if p.Temperature < 0 {
		p.Temperature = p.Temperature * -1
	}
	return p
}
