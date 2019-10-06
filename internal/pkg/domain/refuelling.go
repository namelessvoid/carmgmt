package domain

import "time"

type Refuelling struct {
	ID            int
	CarID         int
	Amount        float32
	Price         float32
	PricePerLiter float32
	Time          time.Time
	Kilometers    float32
}
