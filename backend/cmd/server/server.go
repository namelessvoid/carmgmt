package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/datastore"
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
	for _, pair := range os.Environ() {
		fmt.Println(pair)
	}

	ctx := context.Background()

	logger, err := createLogger()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	config, err := GetConfigFromEnvironment()
	if err != nil {
		panic(err)
	}

	datastoreClient, err := datastore.NewClient(ctx, "carmanagement")
	if err != nil {
		panic(err)
	}
	defer datastoreClient.Close()

	vehicleRepository := domain.NewVehicleRepository()
	vehicleService := domain.NewVehicleService(vehicleRepository, logger)

	authenticator := auth.NewAuthenticator(config.Auth.Issuer, config.Auth.ClientID)

	r := mux.NewRouter()
	mux.CORSMethodMiddleware(r)
	r.Use(auth.AuthenticationMiddleware(authenticator))

	api.ConfigureRoutes(r, vehicleService, logger)

	logger.Info("Running server on " + hostname)
	logger.Fatal(http.ListenAndServe(hostname, api.CORSMiddleware()(r)).Error())
}
