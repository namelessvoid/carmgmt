package domain

import (
	"testing"
)

func Test_NewRefuelling(t *testing.T) {
	tests := []struct {
		name               string
		command            CreateRefuellingCommand
		expectedRefuelling Refuelling
		expectesError      bool
	}{
		{
			name:               "Command is invalid",
			command:            NewCreateRefuellingCommandBuilder().withNilAmount().build(),
			expectedRefuelling: Refuelling{},
			expectesError:      true,
		},
		{
			name:               "Command is valid",
			command:            NewCreateRefuellingCommandBuilder().build(),
			expectedRefuelling: NewRefuellingTestBuilder().Build(),
			expectesError:      false,
		},
	}

	for _, run := range tests {
		t.Run(run.name, func(t *testing.T) {
			r, err := NewRefuelling(run.command)

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
