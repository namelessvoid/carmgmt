package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/namelessvoid/carmgmt/internal/pkg/domain"
)

// GetIndexHandler serves the index page
func getIndexHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Welcome to the car management system!")
}

func getCarByID(cs *domain.CarService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Invalid parameter 'id'", http.StatusBadRequest)
			return
		}

		car, err := cs.GetCarById(id)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		json, err := json.Marshal(car)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	}
}

func ConfigureRoutes(r *mux.Router, cs *domain.CarService) {
	carRouter := r.PathPrefix("/car").Subrouter()
	carRouter.HandleFunc("", getIndexHandler)
	carRouter.HandleFunc("/{id}", getCarByID(cs))
}
