// cars_test.go
// cars_test.go
package cars

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"nextgen-cars/utils"
	"testing"
)

func TestHandleCars(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(HandleCars))
	resp, _ := http.Get(server.URL)

	utils.CheckResponseCode(t, http.StatusOK, resp.StatusCode)

	var respCars []Car
	results, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(results, &respCars)

	if len(respCars) != len(cars) {
		t.Errorf("Expected %d cars, but got %d cars", len(cars), len(respCars))
	}
}

func TestHandleCar(t *testing.T) {
	id := "cjmg09da855l1j6q9240"

	server := httptest.NewServer(http.HandlerFunc(HandleCar))
	resp, _ := http.Get(server.URL + "/car/" + id)
	fmt.Println(server.URL + "/car/" + id)

	utils.CheckResponseCode(t, http.StatusOK, resp.StatusCode)

	var respCar Car
	result, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(result, &respCar)

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
