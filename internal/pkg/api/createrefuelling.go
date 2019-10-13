package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/namelessvoid/carmgmt/internal/pkg/domain"
)

func newCreateRefuellingHandler(vs domain.VehicleService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		// Mux should ensure that vehicleID is present and can be parsed as int.
		// Therefore, a vehicleID which cannot be parsed is an server error due to
		// bad route configuration.
		vehicleID, err := strconv.Atoi(vars["vehicleID"])
		if err != nil {
			internalServerError(w)
			return
		}

		if r.Body == nil {
			invalidJSONError(w)
			return
		}

		cmd := domain.CreateRefuellingCommand{}
		err = json.NewDecoder(r.Body).Decode(&cmd)
		if err != nil {
			invalidJSONError(w)
			return
		}

		cmd.VehicleID = &vehicleID

		refuelling, err := vs.CreateRefuelling(cmd)
		if err != nil {
			internalServerError(w)
			return
		}

		json, err := json.Marshal(refuelling)
		if err != nil {
			internalServerError(w)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	}
}
