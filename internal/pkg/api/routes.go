package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/namelessvoid/carmgmt/internal/pkg/domain"
)

func getAllVehicles(cs *domain.VehicleService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vehicles, err := cs.GetAllVehicles()

		json, err := json.Marshal(vehicles)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	}
}

func createVehicle(cs *domain.VehicleService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var createVehicleCommand domain.Vehicle
		err := json.NewDecoder(r.Body).Decode(&createVehicleCommand)
		if err != nil {
			log.Print(err)
			http.Error(w, "[\"error.invalidJson\"]", http.StatusBadRequest)
			return
		}

		vehicle, err := cs.CreateVehicle(createVehicleCommand.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		json, err := json.Marshal(vehicle)
		if err != nil {
			http.Error(w, "["+err.Error()+"]", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	}
}

func getVehicleByID(cs *domain.VehicleService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Invalid parameter 'id'", http.StatusBadRequest)
			return
		}

		vehicle, err := cs.GetVehicleById(id)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		json, err := json.Marshal(vehicle)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	}
}

func ConfigureRoutes(r *mux.Router, cs *domain.VehicleService) {
	vehicleRouter := r.PathPrefix("/vehicles").Subrouter()
	vehicleRouter.HandleFunc("", getAllVehicles(cs)).Methods("GET")
	vehicleRouter.HandleFunc("", createVehicle(cs)).Methods("POST")
	vehicleRouter.HandleFunc("/{id}", getVehicleByID(cs)).Methods("GET")
}
