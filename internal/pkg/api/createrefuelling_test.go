package api

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/mux"

	"github.com/golang/mock/gomock"
	"github.com/namelessvoid/carmgmt/internal/pkg/domain"
	mock_domain "github.com/namelessvoid/carmgmt/internal/pkg/domain/mocks"
)

func TestCreateRefuellingHandler(t *testing.T) {
	intPtr := func(i int) *int { return &i }
	floatPtr := func(f float32) *float32 { return &f }
	timePtr := func(t time.Time) *time.Time { return &t }

	command := domain.CreateRefuellingCommand{
		VehicleID:     intPtr(200),
		Amount:        floatPtr(30.5),
		Price:         floatPtr(60.47),
		PricePerLiter: floatPtr(1.67),
		Time:          timePtr(time.Date(1994, 11, 5, 13, 15, 30, 0, time.UTC)),
		Kilometers:    floatPtr(823.12),
	}

	tests := []struct {
		name                  string
		requestBody           io.Reader
		expectServiceCall     bool
		expectedCreateCommand domain.CreateRefuellingCommand
		refuellingFromService domain.Refuelling
		errorFromService      error
		expectedResponseCode  int
		expectedResponseBody  string
	}{{
		//
		name:                 "Request body is empty",
		requestBody:          nil,
		expectedResponseCode: http.StatusBadRequest,
		expectedResponseBody: "[\"error.invalidJson\"]"}, {
		//
		name:                 "Request body contains invalid json",
		requestBody:          strings.NewReader("adsfa"),
		expectedResponseCode: http.StatusBadRequest,
		expectedResponseBody: "[\"error.invalidJson\"]"}, {
		//
		name:                  "VehicleService returns error",
		requestBody:           strings.NewReader("{\"amount\":30.5,\"price\":60.47,\"pricePerLiter\":1.67,\"time\":\"1994-11-05T13:15:30Z\",\"kilometers\":823.12}"),
		expectServiceCall:     true,
		expectedCreateCommand: command,
		refuellingFromService: domain.Refuelling{},
		errorFromService:      errors.New("Some internal error"),
		expectedResponseCode:  http.StatusInternalServerError,
		expectedResponseBody:  "[\"error.unknown\"]"}, {
		//
		name:                  "Refuelling created successfully",
		requestBody:           strings.NewReader("{\"amount\":30.5,\"price\":60.47,\"pricePerLiter\":1.67,\"time\":\"1994-11-05T13:15:30Z\",\"kilometers\":823.12}"),
		expectServiceCall:     true,
		expectedCreateCommand: command,
		refuellingFromService: domain.Refuelling{ID: 10, VehicleID: 200, Amount: 30.5, Price: 60.47, PricePerLiter: 1.67, Time: time.Date(1994, 11, 5, 13, 15, 30, 0, time.UTC), Kilometers: 823.12, Consumption: 12.12},
		errorFromService:      nil,
		expectedResponseCode:  http.StatusOK,
		expectedResponseBody:  "{\"id\":10,\"vehicleId\":200,\"amount\":30.5,\"price\":60.47,\"pricePerLiter\":1.67,\"time\":\"1994-11-05T13:15:30Z\",\"kilometers\":823.12,\"consumption\":12.12}"},
	}
	for _, testCfg := range tests {
		t.Run(testCfg.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			vs := mock_domain.NewMockVehicleService(mockCtrl)
			if testCfg.expectServiceCall {
				vs.EXPECT().CreateRefuelling(testCfg.expectedCreateCommand).Return(testCfg.refuellingFromService, testCfg.errorFromService)
			}

			handler := http.HandlerFunc(newCreateRefuellingHandler(vs))
			response := httptest.NewRecorder()

			req, err := http.NewRequest("", "", testCfg.requestBody)
			req = mux.SetURLVars(req, map[string]string{"vehicleID": "200"}) // No need to test, see comment in handler
			if err != nil {
				t.Fatal(err)
			}

			handler.ServeHTTP(response, req)

			if status := response.Code; status != testCfg.expectedResponseCode {
				t.Errorf("create handler returned wrong status code: got %v want %v", status, testCfg.expectedResponseCode)
			}

			if body := response.Body.String(); body != testCfg.expectedResponseBody {
				t.Errorf("create handler returned wrong content: got %v want %v", body, testCfg.expectedResponseBody)
			}

			expectedContentType := "application/json"
			if contentType := response.Header().Get("Content-Type"); contentType != expectedContentType {
				t.Errorf("create handler returned wrong Content-Type: got %v want %v", contentType, expectedContentType)
			}
		})
	}
}
