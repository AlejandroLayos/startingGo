package conections

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/AlejandroLayos/startingGo/src/model"
)

//GetListOfPlanets gets all the planets registered
func GetListOfPlanets() []model.Planet {
	var planets []model.Planet
	spaceClient := http.Client{
		Timeout: time.Second * 10, // Maximum of 10 secs
	}
	var urlPage string = "https://swapi.co/api/planets/?page=1"
	for urlPage != "" {
		var request model.RequestPlanet

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
			planets = append(planets, element)
		}
		urlPage = request.Next

	}
	return planets
}
