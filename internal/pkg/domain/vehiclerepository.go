package domain

import "fmt"

type VehicleRepository interface {
	CreateVehicle(vehicle Vehicle) (Vehicle, error)
	GetAllVehicles() ([]Vehicle, error)
	GetVehicleByID(vehicleID int) (Vehicle, error)
	CreateRefuelling(refuelling Refuelling) (Refuelling, error)
	GetRefuellingsByVehicleID(vehicleID int) ([]Refuelling, error)
}

type vehicleRepository struct {
	vehicles    []Vehicle
	refuellings []Refuelling
}

func NewVehicleRepository() *vehicleRepository {
	return &vehicleRepository{vehicles: []Vehicle{}, refuellings: []Refuelling{}}
}

func (repo *vehicleRepository) doesVehicleExist(vehicleID int) bool {
	return vehicleID >= 0 && vehicleID < len(repo.vehicles)
}

func (repo *vehicleRepository) CreateVehicle(vehicle Vehicle) (Vehicle, error) {
	vehicle.ID = len(repo.vehicles)
	repo.vehicles = append(repo.vehicles, vehicle)
	return vehicle, nil
}

func (repo *vehicleRepository) GetAllVehicles() ([]Vehicle, error) {
	return repo.vehicles, nil
}

func (repo *vehicleRepository) GetVehicleByID(vehicleID int) (Vehicle, error) {
	if !repo.doesVehicleExist(vehicleID) {
		return Vehicle{}, fmt.Errorf("Vehicle with id '%d' does not exist", vehicleID)
	}

	return repo.vehicles[vehicleID], nil
}

func (repo *vehicleRepository) CreateRefuelling(refuelling Refuelling) (Refuelling, error) {
	refuelling.ID = len(repo.refuellings)
	repo.refuellings = append(repo.refuellings, refuelling)
	return refuelling, nil
}

func (repo *vehicleRepository) GetRefuellingsByVehicleID(vehicleID int) ([]Refuelling, error) {
	if !repo.doesVehicleExist(vehicleID) {
		return nil, fmt.Errorf("Vehicle with id '%d' does not exist", vehicleID)
	}

	vehicleRefuellings := []Refuelling{}
	for _, refuelling := range repo.refuellings {
		if refuelling.VehicleID == vehicleID {
			vehicleRefuellings = append(vehicleRefuellings, refuelling)
		}
	}
	return vehicleRefuellings, nil
}
