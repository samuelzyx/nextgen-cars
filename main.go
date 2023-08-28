package main

import (
	"log"
	"net/http"

	"nextgen-cars/cars"
)

func main() {
	http.HandleFunc("/cars", cars.HandleCars)
	http.HandleFunc("/car/", cars.HandleCar)
	log.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
