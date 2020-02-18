package conections

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/AlejandroLayos/startingGo/src/model"
)

//GetListOfStarships gets all the Starships registered
func GetListOfStarships() []model.Starships {
	var starships []model.Starships
	spaceClient := http.Client{
		Timeout: time.Second * 5, // Maximum of 5 secs
	}
	var urlPage string = "https://swapi.co/api/starships/?page=1"
	for urlPage != "" {
		var request model.RequestStarship

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
			starships = append(starships, element)
		}
		urlPage = request.Next

	}
	return starships
}
