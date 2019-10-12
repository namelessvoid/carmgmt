package api

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/gorilla/mux"

	"go.uber.org/zap"

	"github.com/namelessvoid/carmgmt/internal/pkg/domain"
	mock_domain "github.com/namelessvoid/carmgmt/internal/pkg/domain/mocks"

	"github.com/golang/mock/gomock"
)

func Test_getRefuellingsByVehicle(t *testing.T) {
	tests := []struct {
		name                   string
		requestVehicleID       string
		refuellingsFromService []domain.Refuelling
		errorFromService       error
		expectedResponseCode   int
		expectedResponseBody   string
	}{{
		//
		name:                   "Empty refuelling array returned from VehicleService",
		requestVehicleID:       "20",
		refuellingsFromService: []domain.Refuelling{},
		errorFromService:       nil,
		expectedResponseCode:   http.StatusOK,
		expectedResponseBody:   "[]"}, {
		//
		name:                   "Refuellings are returned from VehicleService",
		requestVehicleID:       "20",
		refuellingsFromService: []domain.Refuelling{{ID: 1, VehicleID: 20, Amount: 30.0, Price: 40.0, PricePerLiter: 50.0, Time: time.Date(1994, 3, 12, 13, 14, 15, 0, time.UTC), Kilometers: 70}, {ID: 2, VehicleID: 20, Amount: 33.0, Price: 44.0, PricePerLiter: 55.0, Time: time.Date(1996, 3, 12, 13, 14, 15, 0, time.UTC), Kilometers: 77}},
		errorFromService:       nil,
		expectedResponseCode:   http.StatusOK,
		expectedResponseBody:   "[{\"id\":1,\"vehicleId\":20,\"amount\":30,\"price\":40,\"pricePerLiter\":50,\"time\":\"1994-03-12T13:14:15Z\",\"kilometers\":70},{\"id\":2,\"vehicleId\":20,\"amount\":33,\"price\":44,\"pricePerLiter\":55,\"time\":\"1996-03-12T13:14:15Z\",\"kilometers\":77}]"}, {
		//
		name:                   "VehicleService returns error",
		requestVehicleID:       "20",
		refuellingsFromService: nil,
		errorFromService:       errors.New("Some internal error description"),
		expectedResponseCode:   http.StatusInternalServerError,
		expectedResponseBody:   "[\"error.unknown\"]"}, {
		//
		name:                 "URL parameter vehicleID is not a number",
		requestVehicleID:     "adsadsf",
		expectedResponseCode: http.StatusInternalServerError,
		expectedResponseBody: "[\"error.unknown\"]"}}
	for _, testCfg := range tests {
		t.Run(testCfg.name, func(t *testing.T) {
			logger, _ := zap.NewDevelopment()

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			vs := mock_domain.NewMockVehicleService(mockCtrl)
			vehicleID, err := strconv.Atoi(testCfg.requestVehicleID)
			if err == nil {
				vs.EXPECT().GetRefuellingsByVehicle(vehicleID).Return(testCfg.refuellingsFromService, testCfg.errorFromService)
			}

			handler := http.HandlerFunc(newGetRefuellingsByVehicleHandler(vs, logger))
			response := httptest.NewRecorder()

			req, err := http.NewRequest("", "", nil)
			if err != nil {
				t.Fatal(err)
			}
			req = mux.SetURLVars(req, map[string]string{"vehicleID": testCfg.requestVehicleID})

			handler.ServeHTTP(response, req)

			if status := response.Code; status != testCfg.expectedResponseCode {
				t.Errorf("getRefuellingsByVehicle() returned wrong status code: got %v want %v", status, testCfg.expectedResponseCode)
			}

			if content := response.Body.String(); content != testCfg.expectedResponseBody {
				t.Errorf("getRefuellingsByVehicle() returned wrong content: got %v want %v", content, testCfg.expectedResponseBody)
			}

			expectedContentType := "application/json"
			if contentType := response.Header().Get("Content-Type"); contentType != expectedContentType {
				t.Errorf("getRefuellingsByVehicle() returned wrong Content-Type header: got %v want %v", contentType, expectedContentType)
			}
		})
	}
}
