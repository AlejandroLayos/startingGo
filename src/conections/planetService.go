package conections

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/AlejandroLayos/startingGo/src/model"
)

var urlPlanets string = "https://swapi.co/api/planets/"

//GetListOfPlanets gets all the planets registered
func GetListOfPlanets() []model.Planet {
	var request model.RequestPlanet
	var planets []model.Planet
	spaceClient := http.Client{
		Timeout: time.Second * 5, // Maximum of 10 secs
	}

	req, err := http.NewRequest(http.MethodGet, urlPlanets, nil)
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
	planets = request.Results
	return planets
}

//GetPlanetByURL gets one planet depending of the url
func GetPlanetByURL(url string) model.Planet {
	var planet model.Planet
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
	jsonErr := json.Unmarshal(body, &planet)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return planet
}
