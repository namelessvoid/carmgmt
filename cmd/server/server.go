package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/namelessvoid/carmgmt/internal/pkg/api"

	"github.com/namelessvoid/carmgmt/internal/pkg/domain"

	"github.com/gorilla/mux"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	hostname = ":8080"
)

func createLogger() (*zap.Logger, error) {
	rawJSON := []byte(`{
		"level": "debug",
		"encoding": "json",
		"outputPaths": ["stdout"],
		"errorOutputPaths": ["stderr"],
		"encoderConfig": {
		  "messageKey": "message",
		  "timeKey": "time",
		  "levelKey": "level",
		  "levelEncoder": "lowercase"
		}
	  }`)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoder(func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.UTC().Format("2006-01-02T15:04:05Z0700"))
	})

	return cfg.Build()
}

func main() {
	logger, err := createLogger()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	vehicleService := domain.NewVehicleService(logger)

	vehicle, _ := vehicleService.CreateVehicle("Foo Vehicle")
	_, _ = vehicleService.CreateVehicle("Bar")
	_, err = vehicleService.CreateRefuelling(domain.Refuelling{VehicleID: vehicle.ID})
	if err != nil {
		logger.Error(err.Error())
	}
	_, err = vehicleService.CreateRefuelling(domain.Refuelling{VehicleID: vehicle.ID})
	if err != nil {
		logger.Error(err.Error())
	}
	refuellings, err := vehicleService.GetRefuellingsByVehicle(vehicle.ID)
	if err != nil {
		logger.Error(err.Error())
	}

	fmt.Println(refuellings)

	r := mux.NewRouter()

	api.ConfigureRoutes(r, vehicleService, logger)

	logger.Info("Running server on " + hostname)
	logger.Fatal(http.ListenAndServe(hostname, api.CORSMiddleware()(r)).Error())
}
