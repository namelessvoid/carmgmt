package domain

import (
	"fmt"
	"log"
)

type VehicleService struct{}

var vehicles []Vehicle
var refuellings []Refuelling

func NewVehicleService() *VehicleService {
	return &VehicleService{}
}

func (*VehicleService) CreateVehicle(name string) (Vehicle, error) {
	id := len(vehicles)
	vehicle := Vehicle{ID: id, Name: name}
	vehicles = append(vehicles, vehicle)
	log.Printf("Created vehicle. Number of vehicles: '%d", len(vehicles))
	return vehicle, nil
}

func (*VehicleService) GetAllVehicles() ([]Vehicle, error) {
	return vehicles, nil
}

func (*VehicleService) GetVehicleById(id int) (Vehicle, error) {
	if !doesVehicleExist(id) {
		return Vehicle{ID: -1}, fmt.Errorf("No vehicle with id '%d'", id)
	}

	return vehicles[id], nil
}

func (*VehicleService) AddRefuellingToVehicle(r Refuelling) error {
	if !doesVehicleExist(r.VehicleID) {
		return fmt.Errorf("No vehicle with id '%d'", r.VehicleID)
	}

	refuellings = append(refuellings, r)
	log.Printf("Added refuelling to vehicle with id '%d'. Total number of refuelings: %d", r.VehicleID, len(refuellings))

	return nil
}

func (*VehicleService) GetRefuellingsByVehicle(vehicleID int) ([]Refuelling, error) {
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
