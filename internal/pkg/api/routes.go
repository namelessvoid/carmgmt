package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/namelessvoid/carmgmt/internal/pkg/domain"
)

func getAllCars(cs *domain.CarService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cars, err := cs.GetAllCars()

		json, err := json.Marshal(cars)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	}
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
	carRouter := r.PathPrefix("/cars").Subrouter()
	carRouter.HandleFunc("", getAllCars(cs))
	carRouter.HandleFunc("/{id}", getCarByID(cs))
}
