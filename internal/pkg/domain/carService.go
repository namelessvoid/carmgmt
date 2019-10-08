package domain

import (
	"fmt"
	"log"
)

type CarService struct{}

var cars []Car
var refuellings []Refuelling

func NewCarService() *CarService {
	return &CarService{}
}

func (*CarService) CreateCar(name string) (Car, error) {
	id := len(cars)
	car := Car{ID: id, Name: name}
	cars = append(cars, car)
	log.Printf("Created car. Number of cars: '%d", len(cars))
	return car, nil
}

func (*CarService) GetAllCars() ([]Car, error) {
	return cars, nil
}

func (*CarService) GetCarById(id int) (Car, error) {
	if !doesCarExist(id) {
		return Car{ID: -1}, fmt.Errorf("No car with id '%d'", id)
	}

	return cars[id], nil
}

func (*CarService) AddRefuellingToCar(r Refuelling) error {
	if !doesCarExist(r.CarID) {
		return fmt.Errorf("No car with id '%d'", r.CarID)
	}

	refuellings = append(refuellings, r)
	log.Printf("Added refuelling to car with id '%d'. Total number of refuelings: %d", r.CarID, len(refuellings))

	return nil
}

func (*CarService) GetRefuellingsByCar(carID int) ([]Refuelling, error) {
	if !doesCarExist(carID) {
		return nil, fmt.Errorf("No car with id '%d'", carID)
	}

	var carRefuellings []Refuelling

	for _, r := range refuellings {
		if r.CarID == carID {
			carRefuellings = append(carRefuellings, r)
		}
	}

	return carRefuellings, nil
}

func doesCarExist(id int) bool {
	return id >= 0 && id < len(cars)
}
