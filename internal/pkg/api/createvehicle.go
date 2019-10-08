package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/namelessvoid/carmgmt/internal/pkg/domain"
)

func createVehicle(vs domain.VehicleService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var createVehicleCommand domain.Vehicle
		err := json.NewDecoder(r.Body).Decode(&createVehicleCommand)
		if err != nil {
			log.Print(err)
			http.Error(w, "[\"error.invalidJson\"]", http.StatusBadRequest)
			return
		}

		vehicle, err := vs.CreateVehicle(createVehicleCommand.Name)
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
