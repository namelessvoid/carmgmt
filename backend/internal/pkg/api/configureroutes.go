package api

import (
	"github.com/gorilla/mux"
	"github.com/namelessvoid/carmgmt/internal/pkg/auth"
	"github.com/namelessvoid/carmgmt/internal/pkg/domain"
	"go.uber.org/zap"
)

// ConfigureRoutes adds routes and handlers for the vehicle API to the provided router
func ConfigureRoutes(r *mux.Router, vs domain.VehicleService, logger *zap.Logger) {
	vehicleRouter := r.PathPrefix("/vehicles").Subrouter()
	vehicleRouter.Use(auth.AuthorizationMiddleware(auth.IsAuthenticated))

	vehicleRouter.HandleFunc("", getAllVehicles(vs)).Methods("GET", "OPTIONS")
	vehicleRouter.HandleFunc("", createVehicle(vs)).Methods("POST")
	vehicleRouter.HandleFunc("/{vehicleID:[0-9]+}", getVehicleByID(vs)).Methods("GET")

	vehicleRouter.HandleFunc("/{vehicleID:[0-9]+}/refuellings", newGetRefuellingsByVehicleHandler(vs, logger)).Methods("GET")
	vehicleRouter.HandleFunc("/{vehicleID:[0-9]+}/refuellings", newCreateRefuellingHandler(vs)).Methods("POST")
}
