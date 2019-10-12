package domain

import (
	"fmt"

	"go.uber.org/zap"
)

type VehicleService interface {
	CreateVehicle(name string) (Vehicle, error)
	GetAllVehicles() ([]Vehicle, error)
	GetVehicleByID(id int) (Vehicle, error)
	CreateRefuelling(r Refuelling) (Refuelling, error)
	GetRefuellingsByVehicle(vehicleID int) ([]Refuelling, error)
}

type vehicleService struct {
	logger *zap.Logger
}

var vehicles []Vehicle
var refuellings []Refuelling

func NewVehicleService(l *zap.Logger) *vehicleService {
	logger := l
	if logger == nil {
		logger = zap.NewNop()
	}

	return &vehicleService{logger: logger}
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

func (*vehicleService) GetVehicleByID(id int) (Vehicle, error) {
	if !doesVehicleExist(id) {
		return Vehicle{ID: -1}, fmt.Errorf("No vehicle with id '%d'", id)
	}

	return vehicles[id], nil
}

func (vs *vehicleService) CreateRefuelling(r Refuelling) (Refuelling, error) {
	if !doesVehicleExist(r.VehicleID) {
		return Refuelling{}, fmt.Errorf("No vehicle with id '%d'", r.VehicleID)
	}

	r.ID = len(refuellings)
	refuellings = append(refuellings, r)
	vs.logger.Info("Added refuelling to vehicle.", zap.Int("vehicleId", r.VehicleID), zap.Int("totalRefulingsCount", len(refuellings)))

	return r, nil
}

func (*vehicleService) GetRefuellingsByVehicle(vehicleID int) ([]Refuelling, error) {
	if !doesVehicleExist(vehicleID) {
		return nil, fmt.Errorf("No vehicle with id '%d'", vehicleID)
	}

	vehicleRefuellings := []Refuelling{}

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
