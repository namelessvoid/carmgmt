package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/namelessvoid/carmgmt/internal/pkg/api"

	"github.com/namelessvoid/carmgmt/internal/pkg/domain"

	"github.com/gorilla/mux"
)

const (
	hostname = ":8080"
)

func main() {
	vehicleService := domain.NewVehicleService()
	vehicle, _ := vehicleService.CreateVehicle("Foo Vehicle")
	_, _ = vehicleService.CreateVehicle("Bar")
	err := vehicleService.AddRefuellingToVehicle(domain.Refuelling{VehicleID: vehicle.ID})
	if err != nil {
		log.Println(err)
	}
	err = vehicleService.AddRefuellingToVehicle(domain.Refuelling{VehicleID: vehicle.ID})
	if err != nil {
		log.Println(err)
	}
	refuellings, err := vehicleService.GetRefuellingsByVehicle(vehicle.ID)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(refuellings)

	r := mux.NewRouter()

	api.ConfigureRoutes(r, vehicleService)

	log.Printf("Running server on %s\n", hostname)
	log.Fatal(http.ListenAndServe(hostname, api.CORSMiddleware()(r)))
}
