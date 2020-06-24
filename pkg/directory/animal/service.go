package animal

import (
	"errors"

	"github.com/iyu/go-example/pkg/directory/animal/category"
)

type Service interface {
	GetName(carCategory category.AnimalCategory) (string, error)
	Move(carCategory category.AnimalCategory) (int, error)
}

type service struct {
	categories map[category.AnimalCategory]category.Animal
}

func New(animals category.AnimalSlice) Service {
	s := &service{make(map[category.AnimalCategory]category.Animal)}
	for _, car := range animals {
		s.setCategory(car)
	}
	return s
}

func (s *service) setCategory(car category.Animal) {
	s.categories[car.GetCategory()] = car
}

func (s *service) getCategory(carCategory category.AnimalCategory) (category.Animal, error) {
	car, ok := s.categories[carCategory]
	if !ok {
		return nil, errors.New("not found")
	}
	return car, nil
}

func (s *service) GetName(carCategory category.AnimalCategory) (string, error) {
	car, err := s.getCategory(carCategory)
	if err != nil {
		return "", err
	}
	return car.GetName(), nil
}

func (s *service) Move(carCategory category.AnimalCategory) (int, error) {
	car, err := s.getCategory(carCategory)
	if err != nil {
		return 0, err
	}
	return car.Move(), nil
}

