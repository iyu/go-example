package category

type cat struct {
	name string
	mobility int
}

func NewCat() Animal {
	return &cat{"cat", 3}
}

func (c *cat) GetCategory() AnimalCategory {
	return Cat
}

func (c *cat) GetName() string {
	return c.name
}

func (c *cat) Move() int {
	return c.mobility
}
