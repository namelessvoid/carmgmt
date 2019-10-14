package api

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/namelessvoid/carmgmt/internal/pkg/domain"
	domain_mock "github.com/namelessvoid/carmgmt/internal/pkg/domain/mock"
)

func TestCreateVehicle(t *testing.T) {
	tests := []struct {
		name                 string
		requestBody          io.Reader
		serviceIsCalled      bool
		vehicleName          string
		vehicleFromService   domain.Vehicle
		errorFromService     error
		expectedResponseCode int
		expectedResponseBody string
	}{{
		//
		name:                 "Request body is empty",
		requestBody:          nil,
		expectedResponseCode: http.StatusBadRequest,
		expectedResponseBody: "[\"error.invalidJson\"]"}, {
		//
		name:                 "Request body contains invalid json",
		requestBody:          strings.NewReader("invalidjson"),
		expectedResponseCode: http.StatusBadRequest,
		expectedResponseBody: "[\"error.invalidJson\"]"}, {
		//
		name:                 "Vehicle created successfully",
		requestBody:          strings.NewReader("{\"name\":\"VW Polo\"}"),
		serviceIsCalled:      true,
		vehicleName:          "VW Polo",
		vehicleFromService:   domain.Vehicle{ID: 212, Name: "VW Polo"},
		errorFromService:     nil,
		expectedResponseCode: http.StatusOK,
		expectedResponseBody: "{\"id\":212,\"name\":\"VW Polo\"}"}, {
		//
		name:                 "VehicleService returns error",
		requestBody:          strings.NewReader("{\"name\":\"VW Polo\"}"),
		serviceIsCalled:      true,
		vehicleName:          "VW Polo",
		errorFromService:     errors.New("Some internal error"),
		expectedResponseCode: http.StatusInternalServerError,
		expectedResponseBody: "[\"error.unknown\"]"}}

	for _, testCfg := range tests {
		t.Run(testCfg.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			vs := domain_mock.NewMockVehicleService(mockCtrl)
			if testCfg.serviceIsCalled {
				vs.EXPECT().CreateVehicle(testCfg.vehicleName).Return(testCfg.vehicleFromService, testCfg.errorFromService)
			}

			handler := http.HandlerFunc(createVehicle(vs))
			response := httptest.NewRecorder()

			req, err := http.NewRequest("", "", testCfg.requestBody)
			if err != nil {
				t.Fatal(err)
			}

			handler.ServeHTTP(response, req)

			if status := response.Code; status != testCfg.expectedResponseCode {
				t.Errorf("createVehicle() returned wrong status code: got %v want %v", status, testCfg.expectedResponseCode)
			}

			if body := response.Body.String(); body != testCfg.expectedResponseBody {
				t.Errorf("createVehicle() returned wrong content: got %v want %v", body, testCfg.expectedResponseBody)
			}

			expectedContentType := "application/json"
			if contentType := response.Header().Get("Content-Type"); contentType != expectedContentType {
				t.Errorf("createVehicle() returned wrong Content-Type: got %v want %v", contentType, expectedContentType)
			}
		})
	}
}
