package conections

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/AlejandroLayos/startingGo/src/model"
)

//GetListOfSpecies gets all the Species registered
func GetListOfSpecies() []model.Species {
	var species []model.Species
	spaceClient := http.Client{
		Timeout: time.Second * 5, // Maximum of 10 secs
	}
	var urlPage string = "https://swapi.co/api/species/?page=1"
	for urlPage != "" {
		var request model.RequestSpecies

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
			species = append(species, element)
		}
		urlPage = request.Next

	}
	return species
}
