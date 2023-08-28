package cars

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net/http"
)

// Car Struct
type Car struct {
	ID       string `json:"id"`
	Make     string `json:"make"`
	Model    string `json:"model"`
	Package  string `json:"package"`
	Color    string `json:"color"`
	Year     int    `json:"year"`
	Category string `json:"category"`
	Milage   int    `json:"milage"`
	Price    int    `json:"price"`
}

var cars = []Car{
	{ID: "JHk290Xj", Make: "Ford", Model: "F10", Package: "Base", Color: "Silver", Year: 2010, Category: "Truck", Milage: 120123, Price: 1999990},
	{ID: "fWI37la", Make: "Toyota", Model: "Camry", Package: "SE", Color: "White", Year: 2019, Category: "Sedan", Milage: 3999, Price: 2899000},
	{ID: "1!3xjRllC", Make: "Toyota", Model: "Rav4", Package: "XSE", Color: "Red", Year: 2018, Category: "SUV", Milage: 24001, Price: 2275000},
	{ID: "dku43920s", Make: "Ford", Model: "Bronco", Package: "Bandlands", Color: "Burnt Orange", Year: 2022, Category: "SUV", Milage: 1, Price: 4499000},
}

func handleCars(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cars)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleCar(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/car/"):]
	switch r.Method {
	case http.MethodGet:
		car := findCarByID(id)
		if car == nil {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(car)
	case http.MethodPost:
		var newCar Car
		err := json.NewDecoder(r.Body).Decode(&newCar)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		newCar.ID = generateID()
		cars = append(cars, newCar)
		w.WriteHeader(http.StatusCreated)
	case http.MethodPut:
		var updatedCar Car
		err := json.NewDecoder(r.Body).Decode(&updatedCar)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		index := findCarIndexByID(id)
		if index == -1 {
			http.NotFound(w, r)
			return
		}
		updatedCar.ID = id
		cars[index] = updatedCar
		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func findCarByID(id string) *Car {
	for _, car := range cars {
		if car.ID == id {
			return &car
		}
	}
	return nil
}

func generateID() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func findCarIndexByID(id string) int {
	for i, car := range cars {
		if car.ID == id {
			return i
		}
	}
	return -1
}
