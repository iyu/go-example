package category

type dog struct {
	name string
	mobility int
}

func NewDog() Animal {
	return &dog{"dog", 10}
}

func (c *dog) GetCategory() AnimalCategory {
	return Dog
}

func (c *dog) GetName() string {
	return c.name
}

func (c *dog) Move() int {
	return c.mobility
}
