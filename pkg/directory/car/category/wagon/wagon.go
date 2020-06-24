package wagon

import (
"github.com/iyu/go-example/pkg/directory/car/category"
)

type wagon struct {
	name string
	mobility int
}

func New() category.Car {
	return &wagon{"wagon", 5}
}

func (c *wagon) GetCategory() category.CarCategory {
	return category.Wagon
}

func (c *wagon) GetName() string {
	return c.name
}

func (c *wagon) Move() int {
	return c.mobility
}
