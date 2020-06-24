package sedan

import (
	"github.com/iyu/go-example/pkg/directory/car/category"
)

type sedan struct {
	name string
	mobility int
}

func New() category.Car {
	return &sedan{"sedan", 10}
}

func (c *sedan) GetCategory() category.CarCategory {
	return category.Sedan
}

func (c *sedan) GetName() string {
	return c.name
}

func (c *sedan) Move() int {
	return c.mobility
}
