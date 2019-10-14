package domain

import "time"

type createRefuellingCommandBuilder struct {
	CreateRefuellingCommand
}

func NewCreateRefuellingCommandBuilder() createRefuellingCommandBuilder {
	intPtr := func(i int) *int { return &i }
	floatPtr := func(f float32) *float32 { return &f }
	timePtr := func(t time.Time) *time.Time { return &t }

	return createRefuellingCommandBuilder{
		CreateRefuellingCommand{
			VehicleID:     intPtr(1),
			Amount:        floatPtr(2.2),
			Price:         floatPtr(3.3),
			PricePerLiter: floatPtr(4.4),
			Time:          timePtr(time.Date(2035, 3, 20, 0, 44, 12, 0, time.UTC)),
			Kilometers:    floatPtr(5.5)},
	}
}

func (cb createRefuellingCommandBuilder) build() CreateRefuellingCommand {
	return cb.CreateRefuellingCommand
}

func (cb createRefuellingCommandBuilder) withNilVehicleID() createRefuellingCommandBuilder {
	cb.VehicleID = nil
	return cb
}

func (cb createRefuellingCommandBuilder) withNilAmount() createRefuellingCommandBuilder {
	cb.Amount = nil
	return cb
}

func (cb createRefuellingCommandBuilder) withAmount(a float32) createRefuellingCommandBuilder {
	cb.Amount = &a
	return cb
}

func (cb createRefuellingCommandBuilder) withNilPrice() createRefuellingCommandBuilder {
	cb.Price = nil
	return cb
}

func (cb createRefuellingCommandBuilder) withPrice(p float32) createRefuellingCommandBuilder {
	cb.Price = &p
	return cb
}

func (cb createRefuellingCommandBuilder) withNilPricePerLiter() createRefuellingCommandBuilder {
	cb.PricePerLiter = nil
	return cb
}

func (cb createRefuellingCommandBuilder) withPricePerLiter(p float32) createRefuellingCommandBuilder {
	cb.PricePerLiter = &p
	return cb
}

func (cb createRefuellingCommandBuilder) withNilTime() createRefuellingCommandBuilder {
	cb.Time = nil
	return cb
}

func (cb createRefuellingCommandBuilder) withTime(t time.Time) createRefuellingCommandBuilder {
	cb.Time = &t
	return cb
}

func (cb createRefuellingCommandBuilder) withNilKilometers() createRefuellingCommandBuilder {
	cb.Kilometers = nil
	return cb
}

func (cb createRefuellingCommandBuilder) withKilometers(p float32) createRefuellingCommandBuilder {
	cb.Kilometers = &p
	return cb
}

type refuellingTestBuilder struct {
	Refuelling
}

func NewRefuellingTestBuilder() refuellingTestBuilder {
	return refuellingTestBuilder{Refuelling{
		VehicleID:     1,
		Amount:        2.2,
		Price:         3.3,
		PricePerLiter: 4.4,
		Time:          time.Date(2035, 3, 20, 0, 44, 12, 0, time.UTC),
		Kilometers:    5.5,
		Consumption:   40}}
}

func (rb refuellingTestBuilder) Build() Refuelling {
	return rb.Refuelling
}
