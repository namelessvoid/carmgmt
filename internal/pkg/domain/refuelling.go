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
}
