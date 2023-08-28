package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/car/list/", cars.handleCars)
	http.HandleFunc("/car/", cars.handleCar)
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
