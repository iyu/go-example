package category

type AnimalCategory int
const (
	Cat AnimalCategory = 1
	Dog AnimalCategory = 2
)

type Animal interface {
	GetName() string
	GetCategory() AnimalCategory
	Move() int
}

type AnimalSlice []Animal
