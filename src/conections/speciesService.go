package conections

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/AlejandroLayos/startingGo/src/model"
)

var urlSpecies string = "https://swapi.co/api/species/"

//GetListOfSpecies gets all the Species registered
func GetListOfSpecies() []model.Species {
	var request model.RequestSpecies
	var species []model.Species
	spaceClient := http.Client{
		Timeout: time.Second * 5, // Maximum of 10 secs
	}

	req, err := http.NewRequest(http.MethodGet, urlSpecies, nil)
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
	species = request.Results
	return species
}

//GetSpeciesByURL gets one Species depending of the url
func GetSpeciesByURL(url string) model.Species {
	var species model.Species
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
	jsonErr := json.Unmarshal(body, &species)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return species
}
