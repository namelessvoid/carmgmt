package api

import (
	"github.com/gorilla/mux"
	"github.com/namelessvoid/carmgmt/internal/pkg/domain"
)

// ConfigureRoutes adds routes and handlers for the vehicle API to the provided router
func ConfigureRoutes(r *mux.Router, vs domain.VehicleService) {
	vehicleRouter := r.PathPrefix("/vehicles").Subrouter()
	vehicleRouter.HandleFunc("", getAllVehicles(vs)).Methods("GET")
	vehicleRouter.HandleFunc("", createVehicle(vs)).Methods("POST")
	vehicleRouter.HandleFunc("/{id}", getVehicleByID(vs)).Methods("GET")
}
