package entity

type Prediction struct {
	UnixTime    int16
	Temperature int8
	Rain        bool
	Snow        bool
	Wind        Wind
	By          string
	//humity, uv, cloudy, thermal sensetion, visibility, fog, pressure
}

func (*Prediction) Make(unixTime int16, temperature int8, rain, snow bool, wind Wind, by string) Prediction {
	return Prediction{unixTime, temperature, rain, snow, wind, by}
}

func (p *Prediction) ForcePositive() *Prediction {
	if p.Temperature < 0 {
		p.Temperature = p.Temperature * -1
	}
	return p
}
