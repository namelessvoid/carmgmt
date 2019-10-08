package domain

import (
	"fmt"
	"log"
)

type VehicleService interface {
	CreateVehicle(name string) (Vehicle, error)
	GetAllVehicles() ([]Vehicle, error)
	GetVehicleById(id int) (Vehicle, error)
}

type vehicleService struct{}

var vehicles []Vehicle
var refuellings []Refuelling

func NewVehicleService() *vehicleService {
	return &vehicleService{}
}

func (*vehicleService) CreateVehicle(name string) (Vehicle, error) {
	id := len(vehicles)
	vehicle := Vehicle{ID: id, Name: name}
	vehicles = append(vehicles, vehicle)
	log.Printf("Created vehicle. Number of vehicles: '%d", len(vehicles))
	return vehicle, nil
}

func (*vehicleService) GetAllVehicles() ([]Vehicle, error) {
	return vehicles, nil
}

func (*vehicleService) GetVehicleById(id int) (Vehicle, error) {
	if !doesVehicleExist(id) {
		return Vehicle{ID: -1}, fmt.Errorf("No vehicle with id '%d'", id)
	}

	return vehicles[id], nil
}

func (*vehicleService) AddRefuellingToVehicle(r Refuelling) error {
	if !doesVehicleExist(r.VehicleID) {
		return fmt.Errorf("No vehicle with id '%d'", r.VehicleID)
	}

	refuellings = append(refuellings, r)
	log.Printf("Added refuelling to vehicle with id '%d'. Total number of refuelings: %d", r.VehicleID, len(refuellings))

	return nil
}

func (*vehicleService) GetRefuellingsByVehicle(vehicleID int) ([]Refuelling, error) {
	if !doesVehicleExist(vehicleID) {
		return nil, fmt.Errorf("No vehicle with id '%d'", vehicleID)
	}

	var vehicleRefuellings []Refuelling

	for _, r := range refuellings {
		if r.VehicleID == vehicleID {
			vehicleRefuellings = append(vehicleRefuellings, r)
		}
	}

	return vehicleRefuellings, nil
}

func doesVehicleExist(id int) bool {
	return id >= 0 && id < len(vehicles)
}
