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

		if r.Body == nil {
			invalidJSONError(w)
			return
		}

		err := json.NewDecoder(r.Body).Decode(&createVehicleCommand)
		if err != nil {
			log.Print(err)
			invalidJSONError(w)
			return
		}

		vehicle, err := vs.CreateVehicle(createVehicleCommand.Name)
		if err != nil {
			log.Print(err)
			internalServerError(w)
			return
		}

		json, err := json.Marshal(vehicle)
		if err != nil {
			log.Print(err)
			internalServerError(w)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	}
}
