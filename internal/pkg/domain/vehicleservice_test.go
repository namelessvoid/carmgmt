package domain_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	domain "github.com/namelessvoid/carmgmt/internal/pkg/domain"
	domain_mock "github.com/namelessvoid/carmgmt/internal/pkg/domain/mock"
)

func Test_VehicleService_CreateRefuelling(t *testing.T) {
	tests := []struct {
		name               string
		command            domain.CreateRefuellingCommand
		expectsRepoCall    bool
		refuellingToRepo   domain.Refuelling
		refuellingFromRepo domain.Refuelling
		errorFromRepo      error
		expectedRefuelling domain.Refuelling
		expectsError       bool
	}{
		{
			name:               "Valid command",
			command:            domain.NewCreateRefuellingCommandBuilder().Build(),
			expectsRepoCall:    true,
			refuellingToRepo:   domain.NewRefuellingTestBuilder().Build(),
			refuellingFromRepo: domain.NewRefuellingTestBuilder().WithID(100).Build(),
			errorFromRepo:      nil,
			expectsError:       false,
		},
		{
			name:            "Command is invalid",
			command:         domain.NewCreateRefuellingCommandBuilder().WithNilAmount().Build(),
			expectsRepoCall: false,
			expectsError:    true,
		},
		{
			name:               "Repository returns error",
			command:            domain.NewCreateRefuellingCommandBuilder().Build(),
			expectsRepoCall:    true,
			refuellingToRepo:   domain.NewRefuellingTestBuilder().Build(),
			refuellingFromRepo: domain.Refuelling{},
			errorFromRepo:      errors.New("Some repo error"),
			expectsError:       true,
		},
	}

	for _, run := range tests {
		t.Run(run.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			vr := domain_mock.NewMockVehicleRepository(mockCtrl)
			if run.expectsRepoCall {
				vr.EXPECT().CreateRefuelling(run.refuellingToRepo).Return(run.refuellingFromRepo, run.errorFromRepo)
			}

			vs := domain.NewVehicleService(vr, nil)

			r, err := vs.CreateRefuelling(run.command)

			if !run.expectsError && r != run.refuellingFromRepo {
				t.Errorf("vehicleService returned incorrect Refuelling: got %v want %v", r, run.refuellingFromRepo)
			}

			if !run.expectsError && err != nil {
				t.Errorf("vehicleService returned unexpected error: got %v want no error", err)
			}

			if run.expectsError && err == nil {
				t.Errorf("vehicleService expected to return error but returned no error")
			}
		})
	}
}
