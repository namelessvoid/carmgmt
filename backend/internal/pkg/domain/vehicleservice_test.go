package domain_test

import (
	"errors"
	"reflect"
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

func Test_VehicleService_GetRefuellingsByVehicle(t *testing.T) {
	t.Run("Repository returns result", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		vehicleID := 12
		expectedRefuellings := []domain.Refuelling{
			domain.NewRefuellingTestBuilder().WithID(13).Build(),
			domain.NewRefuellingTestBuilder().WithID(14).Build()}
		var expectedError error = nil

		refuellingsFromRepo := make([]domain.Refuelling, len(expectedRefuellings))
		copy(refuellingsFromRepo, expectedRefuellings)

		vr := domain_mock.NewMockVehicleRepository(mockCtrl)
		vr.EXPECT().GetRefuellingsByVehicleID(vehicleID).Return(refuellingsFromRepo, expectedError)

		vs := domain.NewVehicleService(vr, nil)

		actualRefuellings, actualError := vs.GetRefuellingsByVehicle(vehicleID)

		if actualError != nil {
			t.Errorf("VehicleService returned unexpected error: got %v want no error", actualError)
		}

		if !reflect.DeepEqual(actualRefuellings, expectedRefuellings) {
			t.Errorf("VehicleService returned incorrect refuellings: got %v want %v", actualRefuellings, expectedRefuellings)
		}
	})

	t.Run("Repository returns error", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		vehicleID := 12

		refuelingsFromRepo := []domain.Refuelling{domain.NewRefuellingTestBuilder().Build()}
		expectedError := errors.New("Some error")

		vr := domain_mock.NewMockVehicleRepository(mockCtrl)
		vr.EXPECT().GetRefuellingsByVehicleID(vehicleID).Return(refuelingsFromRepo, expectedError)

		vs := domain.NewVehicleService(vr, nil)

		actualRefuellings, actualError := vs.GetRefuellingsByVehicle(vehicleID)

		if len(actualRefuellings) != 0 {
			t.Errorf("VehicleService returned non empty vehicle slice: got %v want []", actualRefuellings)
		}

		if actualError != expectedError {
			t.Errorf("VehicleService returned unexpected error: got '%v' want '%v'", actualError, expectedError)
		}
	})
}
