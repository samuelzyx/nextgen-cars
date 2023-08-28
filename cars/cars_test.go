// cars_test.go
// cars_test.go
package cars

import (
	"encoding/json"
	"net/http"
	"nextgen-cars/utils"
	"testing"
)

func TestHandleCars(t *testing.T) {
	req, _ := http.NewRequest("GET", "/cars", nil)
	rec := utils.ExecuteRequest(req)

	utils.CheckResponseCode(t, http.StatusOK, rec.Code)

	var respCars []Car
	json.Unmarshal(rec.Body.Bytes(), &respCars)

	if len(respCars) != len(cars) {
		t.Errorf("Expected %d cars, but got %d cars", len(cars), len(respCars))
	}
}

func TestHandleCar(t *testing.T) {
	id := "cjmg09da855l1j6q9240"
	req, _ := http.NewRequest("GET", "/car/"+id, nil)
	rec := utils.ExecuteRequest(req)

	utils.CheckResponseCode(t, http.StatusOK, rec.Code)

	var respCar Car
	json.Unmarshal(rec.Body.Bytes(), &respCar)

	if respCar.UID != id {
		t.Errorf("Expected car with UID %s, but got car with UID %s", id, respCar.UID)
	}
}

func TestFindCarByUID(t *testing.T) {
	id := "cjmg0fla855kfi0gllt0"
	car := FindCarByUID(id)

	if car == nil {
		t.Errorf("Expected car with UID %s, but got nil", id)
	}
	if car.UID != id {
		t.Errorf("Expected car with UID %s, but got car with UID %s", id, car.UID)
	}
}

func TestFindCarIndexByUID(t *testing.T) {
	id := "cjmg0fla855kfi0gllt0"
	index := FindCarIndexByUID(id)

	if index == -1 {
		t.Errorf("Expected index for car with UID %s, but got -1", id)
	}
}
