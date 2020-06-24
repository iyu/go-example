package category

type CarCategory int
const (
	Sedan CarCategory = 1
	Wagon CarCategory = 2
)

type Car interface {
	GetName() string
	GetCategory() CarCategory
	Move() int
}

type CarSlice []Car
