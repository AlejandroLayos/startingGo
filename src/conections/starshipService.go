package conections

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/AlejandroLayos/startingGo/src/model"
)

var urlStarships string = "https://swapi.co/api/starships/"

//GetListOfStarships gets all the Starships registered
func GetListOfStarships() []model.Starships {
	var request model.RequestStarship
	var starships []model.Starships
	spaceClient := http.Client{
		Timeout: time.Second * 5, // Maximum of 5 secs
	}

	req, err := http.NewRequest(http.MethodGet, urlStarships, nil)
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
	starships = request.Results
	return starships
}

//GetStarshipsByURL gets one Starships depending of the url
func GetStarshipsByURL(url string) model.Starships {
	var starships model.Starships
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
	jsonErr := json.Unmarshal(body, &starships)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return starships
}
