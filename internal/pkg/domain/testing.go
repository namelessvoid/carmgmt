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

func (cb createRefuellingCommandBuilder) Build() CreateRefuellingCommand {
	return cb.CreateRefuellingCommand
}

func (cb createRefuellingCommandBuilder) WithNilVehicleID() createRefuellingCommandBuilder {
	cb.VehicleID = nil
	return cb
}

func (cb createRefuellingCommandBuilder) WithNilAmount() createRefuellingCommandBuilder {
	cb.Amount = nil
	return cb
}

func (cb createRefuellingCommandBuilder) WithAmount(a float32) createRefuellingCommandBuilder {
	cb.Amount = &a
	return cb
}

func (cb createRefuellingCommandBuilder) WithNilPrice() createRefuellingCommandBuilder {
	cb.Price = nil
	return cb
}

func (cb createRefuellingCommandBuilder) WithPrice(p float32) createRefuellingCommandBuilder {
	cb.Price = &p
	return cb
}

func (cb createRefuellingCommandBuilder) WithNilPricePerLiter() createRefuellingCommandBuilder {
	cb.PricePerLiter = nil
	return cb
}

func (cb createRefuellingCommandBuilder) WithPricePerLiter(p float32) createRefuellingCommandBuilder {
	cb.PricePerLiter = &p
	return cb
}

func (cb createRefuellingCommandBuilder) WithNilTime() createRefuellingCommandBuilder {
	cb.Time = nil
	return cb
}

func (cb createRefuellingCommandBuilder) WithTime(t time.Time) createRefuellingCommandBuilder {
	cb.Time = &t
	return cb
}

func (cb createRefuellingCommandBuilder) WithNilKilometers() createRefuellingCommandBuilder {
	cb.Kilometers = nil
	return cb
}

func (cb createRefuellingCommandBuilder) WithKilometers(p float32) createRefuellingCommandBuilder {
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

func (rb refuellingTestBuilder) WithID(id int) refuellingTestBuilder {
	rb.ID = id
	return rb
}

func (rb refuellingTestBuilder) Build() Refuelling {
	return rb.Refuelling
}
