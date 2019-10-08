package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/namelessvoid/carmgmt/internal/pkg/domain"
)

func getAllVehicles(vs domain.VehicleService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vehicles, err := vs.GetAllVehicles()
		if err != nil {
			log.Println(err)
			httpError(w, http.StatusInternalServerError)
			return
		}

		json, err := json.Marshal(vehicles)
		if err != nil {
			log.Println(err)
			httpError(w, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	}
}
