package domain

import (
	"errors"
	"testing"
	"time"
)

func Test_CreateRefuellingCommand_Validate(t *testing.T) {
	tests := []struct {
		name          string
		command       CreateRefuellingCommand
		expectedError error
	}{
		{
			name:          "All parameters are correct",
			command:       NewCreateRefuellingCommandBuilder().build(),
			expectedError: nil,
		}, {
			name:          "VehicleID is null",
			command:       NewCreateRefuellingCommandBuilder().withNilVehicleID().build(),
			expectedError: errors.New("CreateRefuellingCommand.VehicleID must not be null"),
		},
		{
			name:          "Amount is nil",
			command:       NewCreateRefuellingCommandBuilder().withNilAmount().build(),
			expectedError: errors.New("CreateRefuellingCommand.Amount must not be null and must be greater than zero"),
		},
		{
			name:          "Amount is zero",
			command:       NewCreateRefuellingCommandBuilder().withAmount(0).build(),
			expectedError: errors.New("CreateRefuellingCommand.Amount must not be null and must be greater than zero"),
		},
		{
			name:          "Price is nil",
			command:       NewCreateRefuellingCommandBuilder().withNilPrice().build(),
			expectedError: errors.New("CreateRefuellingCommand.Price must not be null and must be greater than zero"),
		},
		{
			name:          "Price is zero",
			command:       NewCreateRefuellingCommandBuilder().withPrice(0).build(),
			expectedError: errors.New("CreateRefuellingCommand.Price must not be null and must be greater than zero"),
		},
		{
			name:          "PricePerLiter is nil",
			command:       NewCreateRefuellingCommandBuilder().withNilPricePerLiter().build(),
			expectedError: errors.New("CreateRefuellingCommand.PricePerLiter must not be null and must be greater than zero"),
		},
		{
			name:          "PricePerLiter is zero",
			command:       NewCreateRefuellingCommandBuilder().withPricePerLiter(0).build(),
			expectedError: errors.New("CreateRefuellingCommand.PricePerLiter must not be null and must be greater than zero"),
		},
		{
			name:          "Time is nil",
			command:       NewCreateRefuellingCommandBuilder().withNilTime().build(),
			expectedError: errors.New("CreateRefuellingCommand.Time must not be null and must be in UTC"),
		},
		{
			name:          "Time is not in UTC",
			command:       NewCreateRefuellingCommandBuilder().withTime(time.Date(2034, 12, 3, 22, 30, 12, 0, &time.Location{})).build(),
			expectedError: errors.New("CreateRefuellingCommand.Time must not be null and must be in UTC"),
		},
		{
			name:          "Kilometers is nil",
			command:       NewCreateRefuellingCommandBuilder().withNilKilometers().build(),
			expectedError: errors.New("CreateRefuellingCommand.Kilometers must not be null and must be greater than zero"),
		},
		{
			name:          "Kilometers is zero",
			command:       NewCreateRefuellingCommandBuilder().withKilometers(0).build(),
			expectedError: errors.New("CreateRefuellingCommand.Kilometers must not be null and must be greater than zero"),
		},
	}

	for _, run := range tests {
		t.Run(run.name, func(t *testing.T) {
			err := run.command.validate()

			if err == nil && run.expectedError != nil {
				t.Errorf("validate() returned no error: want '%v'", run.expectedError)
			}

			if err != nil && run.expectedError == nil {
				t.Errorf("validate() returned unexpected error: got '%v' want '%v'", err, run.expectedError)
			}

			if err != nil && run.expectedError != nil && err.Error() != run.expectedError.Error() {
				t.Errorf("validate() returned unexpected error: got '%v' want '%v'", err, run.expectedError)
			}
		})
	}
}
