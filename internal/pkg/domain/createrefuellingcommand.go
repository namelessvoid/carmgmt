package domain

import (
	"errors"
	"time"
)

type CreateRefuellingCommand struct {
	VehicleID     *int       `json:"vehicleId"`
	Amount        *float32   `json:"amount"`
	Price         *float32   `json:"price"`
	PricePerLiter *float32   `json:"pricePerLiter"`
	Time          *time.Time `json:"time"`
	Kilometers    *float32   `json:"kilometers"`
}

func (c CreateRefuellingCommand) validate() error {
	if c.VehicleID == nil {
		return errors.New("CreateRefuellingCommand.VehicleID must not be null")
	}

	if c.Amount == nil || *c.Amount <= 0 {
		return errors.New("CreateRefuellingCommand.Amount must not be null and must be greater than zero")
	}

	if c.Price == nil || *c.Price <= 0 {
		return errors.New("CreateRefuellingCommand.Price must not be null and must be greater than zero")
	}

	if c.PricePerLiter == nil || *c.PricePerLiter <= 0 {
		return errors.New("CreateRefuellingCommand.PricePerLiter must not be null and must be greater than zero")
	}

	if c.Time == nil || (*c.Time).Location() != time.UTC {
		return errors.New("CreateRefuellingCommand.Time must not be null and must be in UTC")
	}

	if c.Kilometers == nil || *c.Kilometers <= 0 {
		return errors.New("CreateRefuellingCommand.Kilometers must not be null and must be greater than zero")
	}

	return nil
}
