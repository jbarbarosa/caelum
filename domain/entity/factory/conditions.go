package factory

import "caelum/domain/entity"

type ConditionsFactory struct {
	newConditions *entity.Conditions
}

func (c *ConditionsFactory) New() *ConditionsFactory {
	c.newConditions = &entity.Conditions{}
	return c
}

func (c *ConditionsFactory) AddRain() *ConditionsFactory {
	c.newConditions.Rain = true
	return c
}

func (c *ConditionsFactory) AddFog() *ConditionsFactory {
	c.newConditions.Fog = true
	return c
}

func (c *ConditionsFactory) AddSnow() *ConditionsFactory {
	c.newConditions.Snow = true
	return c
}

func (c *ConditionsFactory) AddCloudy() *ConditionsFactory {
	c.newConditions.Cloudy = true
	return c
}

func (c *ConditionsFactory) Finish() entity.Conditions {
	new := c.newConditions
	c.newConditions = nil
	return *new
}
