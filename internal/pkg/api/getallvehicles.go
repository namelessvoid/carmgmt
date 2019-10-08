package api

import (
	"encoding/json"
	"net/http"

	"github.com/namelessvoid/carmgmt/internal/pkg/domain"
)

func getAllVehicles(vs domain.VehicleService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vehicles, err := vs.GetAllVehicles()

		json, err := json.Marshal(vehicles)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	}
}
