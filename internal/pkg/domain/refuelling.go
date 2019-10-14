package domain

import "time"

type Refuelling struct {
	ID            int       `json:"id"`
	VehicleID     int       `json:"vehicleId"`
	Amount        float32   `json:"amount"`
	Price         float32   `json:"price"`
	PricePerLiter float32   `json:"pricePerLiter"`
	Time          time.Time `json:"time"`
	Kilometers    float32   `json:"kilometers"`
	Consumption   float32   `json:"consumption"`
}

func NewRefuelling(cmd CreateRefuellingCommand) (Refuelling, error) {
	err := cmd.Validate()
	if err != nil {
		return Refuelling{}, err
	}

	consumption := 100 * *cmd.Amount / *cmd.Kilometers
	return Refuelling{
		VehicleID:     *cmd.VehicleID,
		Amount:        *cmd.Amount,
		Price:         *cmd.Price,
		PricePerLiter: *cmd.PricePerLiter,
		Time:          *cmd.Time,
		Kilometers:    *cmd.Kilometers,
		Consumption:   consumption}, nil
}
