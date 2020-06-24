package car

import (
	"errors"

	"github.com/iyu/go-example/pkg/directory/car/category"
)

type Service interface {
	GetName(carCategory category.CarCategory) (string, error)
	Move(carCategory category.CarCategory) (int, error)
}

type service struct {
	categories map[category.CarCategory]category.Car
}

func New(cars category.CarSlice) Service {
	s := &service{make(map[category.CarCategory]category.Car)}
	for _, car := range cars {
		s.setCategory(car)
	}
	return s
}

func (s *service) setCategory(car category.Car) {
	s.categories[car.GetCategory()] = car
}

func (s *service) getCategory(carCategory category.CarCategory) (category.Car, error) {
	car, ok := s.categories[carCategory]
	if !ok {
		return nil, errors.New("not found")
	}
	return car, nil
}

func (s *service) GetName(carCategory category.CarCategory) (string, error) {
	car, err := s.getCategory(carCategory)
	if err != nil {
		return "", err
	}
	return car.GetName(), nil
}

func (s *service) Move(carCategory category.CarCategory) (int, error) {
	car, err := s.getCategory(carCategory)
	if err != nil {
		return 0, err
	}
	return car.Move(), nil
}

