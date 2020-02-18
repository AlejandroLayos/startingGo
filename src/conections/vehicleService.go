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
	var vehicles []model.Vehicle
	spaceClient := http.Client{
		Timeout: time.Second * 5, // Maximum of 10 secs
	}
	var urlPage string = "https://swapi.co/api/vehicles/?page=1"

	for urlPage != "" {
		var request model.RequestVehicle

		req, err := http.NewRequest(http.MethodGet, urlPage, nil)
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

		for _, element := range request.Results {
			vehicles = append(vehicles, element)
		}
		urlPage = request.Next

	}
	return vehicles
}
