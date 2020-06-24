package directory

import (
	"github.com/iyu/go-example/pkg/directory/animal"
	animalCategory "github.com/iyu/go-example/pkg/directory/animal/category"
	"github.com/iyu/go-example/pkg/directory/car"
	carCategory "github.com/iyu/go-example/pkg/directory/car/category"
	"github.com/iyu/go-example/pkg/directory/car/category/sedan"
	"github.com/iyu/go-example/pkg/directory/car/category/wagon"
)

func Register() (animal.Service, car.Service) {
	cars := carCategory.CarSlice{sedan.New(), wagon.New()}
	carService := car.New(cars)

	animals := animalCategory.AnimalSlice{animalCategory.NewCat(), animalCategory.NewDog()}
	animalService := animal.New(animals)

	return animalService, carService
}
