package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/namelessvoid/carmgmt/internal/pkg/api"
	"github.com/namelessvoid/carmgmt/internal/pkg/auth"

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

	vehicleRepository := domain.NewVehicleRepository()
	vehicleService := domain.NewVehicleService(vehicleRepository, logger)

	authenticator := auth.NewAuthenticator("username", "password")

	r := mux.NewRouter()
	mux.CORSMethodMiddleware(r)
	r.Use(auth.AuthenticationMiddleware(authenticator))
	r.HandleFunc("/login", auth.LoginHandler(authenticator))

	api.ConfigureRoutes(r, vehicleService, logger)

	logger.Info("Running server on " + hostname)
	logger.Fatal(http.ListenAndServe(hostname, api.CORSMiddleware()(r)).Error())
}
