package factory

import "caelum/domain/entity"

type WindFactory struct {
	newWind *entity.Wind
}

func (w *WindFactory) New() {
	w.newWind = &entity.Wind{}
}

func (w *WindFactory) AddDirection(d entity.Direction) *WindFactory {
	w.newWind.Direction = d
	return w
}

func (w *WindFactory) AddVelocity(velocity int8) *WindFactory {
	w.newWind.Velocity = velocity
	return w
}

func (w *WindFactory) Finish() entity.Wind {
	wind := *w.newWind
	w.newWind = nil
	return wind
}
