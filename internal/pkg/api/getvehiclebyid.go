package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/namelessvoid/carmgmt/internal/pkg/domain"
)

func getVehicleByID(as domain.VehicleService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Invalid parameter 'id'", http.StatusBadRequest)
			return
		}

		vehicle, err := as.GetVehicleById(id)
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
