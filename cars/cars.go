package cars

import (
	"encoding/json"
	"log"
	"net/http"

	"nextgen-cars/utils"
)

// Car Struct
type Car struct {
	UID      string `json:"uid"`
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
	{UID: "cjmg09da855l1j6q9240", ID: "JHk290Xj", Make: "Ford", Model: "F10", Package: "Base", Color: "Silver", Year: 2010, Category: "Truck", Milage: 120123, Price: 1999990},
	{UID: "cjmg0ela855kfi0gllsg", ID: "fWI37la", Make: "Toyota", Model: "Camry", Package: "SE", Color: "White", Year: 2019, Category: "Sedan", Milage: 3999, Price: 2899000},
	{UID: "cjmg0fla855kfi0gllt0", ID: "1!3xjRllC", Make: "Toyota", Model: "Rav4", Package: "XSE", Color: "Red", Year: 2018, Category: "SUV", Milage: 24001, Price: 2275000},
	{UID: "cjmg0e5a855kfi0glls0", ID: "dku43920s", Make: "Ford", Model: "Bronco", Package: "Bandlands", Color: "Burnt Orange", Year: 2022, Category: "SUV", Milage: 1, Price: 4499000},
}

func HandleCars(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		log.Println("Invalid method:", r.Method)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
	log.Println("GET /cars")
}

func HandleCar(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/car/"):]
	switch r.Method {
	case http.MethodGet:
		car := FindCarByUID(id)
		if car == nil {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(car)
		log.Println("GET /car/", id)
	case http.MethodPost:
		var newCar Car
		err := json.NewDecoder(r.Body).Decode(&newCar)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		newCar.ID = utils.GenerateUID()
		cars = append(cars, newCar)
		w.WriteHeader(http.StatusCreated)
		log.Println("POST /car/", newCar.ID)
	case http.MethodPut:
		var updatedCar Car
		err := json.NewDecoder(r.Body).Decode(&updatedCar)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		index := FindCarIndexByUID(id)
		if index == -1 {
			http.NotFound(w, r)
			return
		}
		updatedCar.ID = id
		cars[index] = updatedCar
		w.WriteHeader(http.StatusOK)
		log.Println("PUT /car/", id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func FindCarByUID(id string) *Car {
	for _, car := range cars {
		if car.UID == id {
			return &car
		}
	}
	return nil
}

func FindCarIndexByUID(id string) int {
	for i, car := range cars {
		if car.UID == id {
			return i
		}
	}
	return -1
}
