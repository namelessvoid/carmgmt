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
	carService := domain.NewCarService()
	car := carService.CreateCar("Foo Car")
	_ = carService.CreateCar("Bar")
	err := carService.AddRefuellingToCar(domain.Refuelling{CarID: car.ID})
	if err != nil {
		log.Println(err)
	}
	err = carService.AddRefuellingToCar(domain.Refuelling{CarID: car.ID})
	if err != nil {
		log.Println(err)
	}
	refuellings, err := carService.GetRefuellingsByCar(car.ID)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(refuellings)

	r := mux.NewRouter()

	api.ConfigureRoutes(r, carService)

	log.Printf("Running server on %s\n", hostname)
	log.Fatal(http.ListenAndServe(hostname, api.CORSMiddleware()(r)))
}
