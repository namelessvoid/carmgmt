package domain_test

import (
	"errors"
	"testing"
	"time"

	domain "github.com/namelessvoid/carmgmt/internal/pkg/domain"
)

func Test_CreateRefuellingCommand_Validate(t *testing.T) {
	tests := []struct {
		name          string
		command       domain.CreateRefuellingCommand
		expectedError error
	}{
		{
			name:          "All parameters are correct",
			command:       domain.NewCreateRefuellingCommandBuilder().Build(),
			expectedError: nil,
		}, {
			name:          "VehicleID is null",
			command:       domain.NewCreateRefuellingCommandBuilder().WithNilVehicleID().Build(),
			expectedError: errors.New("CreateRefuellingCommand.VehicleID must not be null"),
		},
		{
			name:          "Amount is nil",
			command:       domain.NewCreateRefuellingCommandBuilder().WithNilAmount().Build(),
			expectedError: errors.New("CreateRefuellingCommand.Amount must not be null and must be greater than zero"),
		},
		{
			name:          "Amount is zero",
			command:       domain.NewCreateRefuellingCommandBuilder().WithAmount(0).Build(),
			expectedError: errors.New("CreateRefuellingCommand.Amount must not be null and must be greater than zero"),
		},
		{
			name:          "Price is nil",
			command:       domain.NewCreateRefuellingCommandBuilder().WithNilPrice().Build(),
			expectedError: errors.New("CreateRefuellingCommand.Price must not be null and must be greater than zero"),
		},
		{
			name:          "Price is zero",
			command:       domain.NewCreateRefuellingCommandBuilder().WithPrice(0).Build(),
			expectedError: errors.New("CreateRefuellingCommand.Price must not be null and must be greater than zero"),
		},
		{
			name:          "PricePerLiter is nil",
			command:       domain.NewCreateRefuellingCommandBuilder().WithNilPricePerLiter().Build(),
			expectedError: errors.New("CreateRefuellingCommand.PricePerLiter must not be null and must be greater than zero"),
		},
		{
			name:          "PricePerLiter is zero",
			command:       domain.NewCreateRefuellingCommandBuilder().WithPricePerLiter(0).Build(),
			expectedError: errors.New("CreateRefuellingCommand.PricePerLiter must not be null and must be greater than zero"),
		},
		{
			name:          "Time is nil",
			command:       domain.NewCreateRefuellingCommandBuilder().WithNilTime().Build(),
			expectedError: errors.New("CreateRefuellingCommand.Time must not be null and must be in UTC"),
		},
		{
			name:          "Time is not in UTC",
			command:       domain.NewCreateRefuellingCommandBuilder().WithTime(time.Date(2034, 12, 3, 22, 30, 12, 0, &time.Location{})).Build(),
			expectedError: errors.New("CreateRefuellingCommand.Time must not be null and must be in UTC"),
		},
		{
			name:          "Kilometers is nil",
			command:       domain.NewCreateRefuellingCommandBuilder().WithNilKilometers().Build(),
			expectedError: errors.New("CreateRefuellingCommand.Kilometers must not be null and must be greater than zero"),
		},
		{
			name:          "Kilometers is zero",
			command:       domain.NewCreateRefuellingCommandBuilder().WithKilometers(0).Build(),
			expectedError: errors.New("CreateRefuellingCommand.Kilometers must not be null and must be greater than zero"),
		},
	}

	for _, run := range tests {
		t.Run(run.name, func(t *testing.T) {
			err := run.command.Validate()

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
