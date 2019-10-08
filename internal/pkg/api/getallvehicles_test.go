package api

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/namelessvoid/carmgmt/internal/pkg/domain"
	mock_domain "github.com/namelessvoid/carmgmt/internal/pkg/domain/mocks"

	"github.com/golang/mock/gomock"
)

func Test_getAllVehicles(t *testing.T) {
	tests := []struct {
		name                 string
		vehiclesFromService  []domain.Vehicle
		errorFromService     error
		expectedResponseCode int
		expectedResponseBody string
	}{{
		//
		name:                 "Empty vehicle array returned from VehicleService",
		vehiclesFromService:  []domain.Vehicle{},
		errorFromService:     nil,
		expectedResponseCode: http.StatusOK,
		expectedResponseBody: "[]"}, {
		//
		name:                 "Vehicles are returned from VehicleService",
		vehiclesFromService:  []domain.Vehicle{{ID: 1, Name: "VW Golf"}, {ID: 200, Name: "Audi A8"}},
		errorFromService:     nil,
		expectedResponseCode: http.StatusOK,
		expectedResponseBody: "[{\"id\":1,\"name\":\"VW Golf\"},{\"id\":200,\"name\":\"Audi A8\"}]"}, {
		//
		name:                 "VehicleService returns error",
		vehiclesFromService:  nil,
		errorFromService:     errors.New("Some internal error description"),
		expectedResponseCode: http.StatusInternalServerError,
		expectedResponseBody: "[\"error.unknown\"]"}}
	for _, testCfg := range tests {
		t.Run(testCfg.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			vs := mock_domain.NewMockVehicleService(ctrl)
			vs.EXPECT().GetAllVehicles().Return(testCfg.vehiclesFromService, testCfg.errorFromService)

			handler := http.HandlerFunc(getAllVehicles(vs))
			response := httptest.NewRecorder()

			req, err := http.NewRequest("POST", "/cars", nil)
			if err != nil {
				t.Fatal(err)
			}

			handler.ServeHTTP(response, req)

			if status := response.Code; status != testCfg.expectedResponseCode {
				t.Errorf("GetAllVehicles returned wrong status code: got %v want %v", status, testCfg.expectedResponseCode)
			}

			if content := response.Body.String(); content != testCfg.expectedResponseBody {
				t.Errorf("GetAllVehicles returned unexpected content: got %v want %v", content, testCfg.expectedResponseBody)
			}

			expectedContentType := "application/json"
			if contentType := response.Header().Get("Content-Type"); contentType != expectedContentType {
				t.Errorf("GetAllVehicles returned unexpected Content-Type header: got %v want %v", contentType, expectedContentType)
			}
		})
	}
}
