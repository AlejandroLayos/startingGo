package conections

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/AlejandroLayos/startingGo/src/model"
)

var urlVechicles string = "https://swapi.co/api/vehicles/"

//GetListOfVehicles gets all the Vehicle registered
func GetListOfVehicles() []model.Vehicle {
	var request model.RequestVehicle
	var vehicles []model.Vehicle
	spaceClient := http.Client{
		Timeout: time.Second * 5, // Maximum of 10 secs
	}

	req, err := http.NewRequest(http.MethodGet, urlVechicles, nil)
	if err != nil {
		log.Fatal(err)
	}
	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	jsonErr := json.Unmarshal(body, &request)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	vehicles = request.Results
	return vehicles
}

//GetVehicleByURL gets one Vehicle depending of the url
func GetVehicleByURL(url string) model.Vehicle {
	var vechicle model.Vehicle
	spaceClient := http.Client{
		Timeout: time.Second * 5, // Maximum of 10 secs
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	jsonErr := json.Unmarshal(body, &vechicle)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return vechicle
}
