package domain

import (
	"fmt"

	"go.uber.org/zap"
)

type VehicleService interface {
	CreateVehicle(name string) (Vehicle, error)
	GetAllVehicles() ([]Vehicle, error)
	GetVehicleById(id int) (Vehicle, error)
}

type vehicleService struct {
	logger *zap.Logger
}

var vehicles []Vehicle
var refuellings []Refuelling

func NewVehicleService(l *zap.Logger) *vehicleService {
	return &vehicleService{logger: l}
}

func (vs *vehicleService) CreateVehicle(name string) (Vehicle, error) {
	id := len(vehicles)
	vehicle := Vehicle{ID: id, Name: name}
	vehicles = append(vehicles, vehicle)
	vs.logger.Info("Created vehicle.", zap.Int("vehicleCount", len(vehicles)))
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

func (vs *vehicleService) AddRefuellingToVehicle(r Refuelling) error {
	if !doesVehicleExist(r.VehicleID) {
		return fmt.Errorf("No vehicle with id '%d'", r.VehicleID)
	}

	refuellings = append(refuellings, r)
	vs.logger.Info("Added refuelling to vehicle.", zap.Int("vehicleId", r.VehicleID), zap.Int("totalRefulingsCount", len(refuellings)))

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
