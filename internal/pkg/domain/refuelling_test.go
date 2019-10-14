package domain_test

import (
	"testing"

	domain "github.com/namelessvoid/carmgmt/internal/pkg/domain"
)

func Test_NewRefuelling(t *testing.T) {
	tests := []struct {
		name               string
		command            domain.CreateRefuellingCommand
		expectedRefuelling domain.Refuelling
		expectesError      bool
	}{
		{
			name:               "Command is invalid",
			command:            domain.NewCreateRefuellingCommandBuilder().WithNilAmount().Build(),
			expectedRefuelling: domain.Refuelling{},
			expectesError:      true,
		},
		{
			name:               "Command is valid",
			command:            domain.NewCreateRefuellingCommandBuilder().Build(),
			expectedRefuelling: domain.NewRefuellingTestBuilder().Build(),
			expectesError:      false,
		},
	}

	for _, run := range tests {
		t.Run(run.name, func(t *testing.T) {
			r, err := domain.NewRefuelling(run.command)

			if r != run.expectedRefuelling {
				t.Errorf("NewRefuelling() returned incorrect Refuelling: got '%v' want '%v'", r, run.expectedRefuelling)
			}

			if err != nil && run.expectesError == false {
				t.Errorf("NewRefuelling() returned unexpected error: got '%v' want no error", err)
			}

			if err == nil && run.expectesError == true {
				t.Error("NewRefuelling() did not return an error but an error was expected")
			}
		})
	}
}
