package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/namelessvoid/carmgmt/internal/pkg/domain"
	"go.uber.org/zap"
)

func newGetRefuellingsByVehicleHandler(vs domain.VehicleService, logger *zap.Logger) http.HandlerFunc {
	if logger == nil {
		logger = zap.NewNop()
	}

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		vehicleID, err := strconv.Atoi(vars["vehicleID"])
		if err != nil {
			logger.Error("InternalServerError due to invalid URL parameter", zap.String("urlParameter", "vehicleID"), zap.String("urlParameterValue", vars["vehicleID"]))
			internalServerError(w)
			return
		}

		refuellings, err := vs.GetRefuellingsByVehicle(vehicleID)
		if err != nil {
			internalServerError(w)
			return
		}

		json, err := json.Marshal(refuellings)
		if err != nil {
			internalServerError(w)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	}
}
