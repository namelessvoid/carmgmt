package domain

import (
	"go.uber.org/zap"
)

type VehicleService interface {
	CreateVehicle(name string) (Vehicle, error)
	GetAllVehicles() ([]Vehicle, error)
	GetVehicleByID(id int) (Vehicle, error)
	CreateRefuelling(cmd CreateRefuellingCommand) (Refuelling, error)
	GetRefuellingsByVehicle(vehicleID int) ([]Refuelling, error)
}

type vehicleService struct {
	logger *zap.Logger
	repo   VehicleRepository
}

var refuellings []Refuelling

func NewVehicleService(repo VehicleRepository, l *zap.Logger) *vehicleService {
	logger := l
	if logger == nil {
		logger = zap.NewNop()
	}

	return &vehicleService{repo: repo, logger: logger}
}

func (vs *vehicleService) CreateVehicle(name string) (Vehicle, error) {
	vehicle := Vehicle{ID: -1, Name: name}
	return vs.repo.CreateVehicle(vehicle)
}

func (vs *vehicleService) GetAllVehicles() ([]Vehicle, error) {
	return vs.repo.GetAllVehicles()
}

func (vs *vehicleService) GetVehicleByID(id int) (Vehicle, error) {
	return vs.repo.GetVehicleByID(id)
}

func (vs *vehicleService) CreateRefuelling(cmd CreateRefuellingCommand) (Refuelling, error) {
	r, err := NewRefuelling(cmd)
	if err != nil {
		return Refuelling{}, err
	}

	return vs.repo.CreateRefuelling(r)
}

func (vs *vehicleService) GetRefuellingsByVehicle(vehicleID int) ([]Refuelling, error) {
	rs, err := vs.repo.GetRefuellingsByVehicleID(vehicleID)
	if err != nil {
		return []Refuelling{}, err
	}

	return rs, nil
}
