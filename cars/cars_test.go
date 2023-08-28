package cars

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

// Importar las librerías necesarias

func TestGetCars(t *testing.T) {
	req, _ := http.NewRequest("GET", "/cars", nil)
	rec := executeRequest(req)

	checkResponseCode(t, http.StatusOK, rec.Code)

	var respCars []Car
	json.Unmarshal(rec.Body.Bytes(), &respCars)

	if len(respCars) != len(cars) {
		t.Errorf("Get: %d, Total: %d", len(cars), len(respCars))
	}
}

func TestGetCar(t *testing.T) {
	id := "JHk290Xj"
	req, _ := http.NewRequest("GET", "/cars/"+id, nil)
	rec := executeRequest(req)

	checkResponseCode(t, http.StatusOK, rec.Code)

	var respCar Car
	json.Unmarshal(rec.Body.Bytes(), &respCar)

	if respCar.ID != id {
		t.Errorf("ID incorrect. Expected: %s, Obtained: %s", id, respCar.ID)
	}
}

func TestPostCar(t *testing.T) {
	newCar := Car{Make: "Honda", Model: "Civic", Package: "EX", Color: "Blue", Year: 2020, Category: "Sedan", Milage: 15000, Price: 2100000}
	payload, _ := json.Marshal(newCar)
	req, _ := http.NewRequest("POST", "/cars", bytes.NewReader(payload))
	rec := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, rec.Code)
}

func TestPutCar(t *testing.T) {
	id := "fWl37la"
	updatedCar := Car{Make: "Toyota", Model: "Camry", Package: "XSE", Color: "Black", Year: 2021, Category: "Sedan", Milage: 5000, Price: 3100000}
	payload, _ := json.Marshal(updatedCar)
	req, _ := http.NewRequest("PUT", "/cars/"+id, bytes.NewReader(payload))
	rec := executeRequest(req)

	checkResponseCode(t, http.StatusOK, rec.Code)

	index := findCarIndexByID(id)
	if index == -1 {
		t.Errorf("Car ID %s not found", id)
	}
	if cars[index].Color != updatedCar.Color {
		t.Errorf("Color not updated. Expected: %s, Obtained: %s", updatedCar.Color, cars[index].Color)
	}
}
