package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/namelessvoid/carmgmt/internal/pkg/domain"
)

type VehicleServiceStub struct {
	*domain.VehicleService
}

func (vstub *VehicleServiceStub) GetAllVehicles() ([]domain.Vehicle, error) {
	return nil, nil
}

func TestGetAllVehicles(t *testing.T) {
	req, err := http.NewRequest("POST", "/cars", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	handler := http.HandlerFunc(getAllVehicles(&VehicleServiceStub{}))
	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
