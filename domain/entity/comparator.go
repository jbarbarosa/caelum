package entity

type Comparator struct {
	Predictions    []Prediction
	CurrentWeather *Prediction
	Error          []string
}

func (c *Comparator) load(p []Prediction) {
	for i := 0; i < len(p); i++ {
		c.Predictions = append(c.Predictions, p[i])
	}
}

func (c *Comparator) setCurrentWeather(p *Prediction) {
	c.CurrentWeather = p
}

func (c *Comparator) run() {
	checkForErrors(c)
	if len(c.Error) > 0 {
		return
	}

}

func (c *Comparator) getBestPrediction(first, second *Prediction) int {
	firstScore := 0
	secondScore := 0
	firstDelta, secondDelta := giveScore(c.getTemperatureDelta(first), c.getTemperatureDelta(second))
	firstScore += firstDelta
	secondScore += secondDelta

	return 1 //todo: finish comparing other struct's attributes and adding to score
}

func (c *Comparator) getTemperatureDelta(p *Prediction) int8 {
	var temperature int8
	if c.CurrentWeather.Temperature < 0 {
		temperature = p.ForcePositive().Temperature - c.CurrentWeather.Temperature
	} else {
		temperature = p.Temperature - c.CurrentWeather.Temperature
	}
	return temperature
}

func giveScore(first, second int8) (int, int) {
	if first < second {
		return 1, 0
	} else {
		return 0, 1
	}
}

func checkForErrors(c *Comparator) {
	if c.Predictions == nil {
		c.Error = append(c.Error, "no predictions loaded")
	}
	if c.CurrentWeather == nil {
		c.Error = append(c.Error, "no current weather to compare")
	}
}
