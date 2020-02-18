package conections

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/AlejandroLayos/startingGo/src/model"
)

//GetListOfCharacters gets all the Characters registered
func GetListOfCharacters() []model.Character {
	var characters []model.Character
	spaceClient := http.Client{
		Timeout: time.Second * 10, // Maximum of 10 secs
	}
	var urlPage string = "https://swapi.co/api/people/?page=1"

	for urlPage != "" {
		var request model.RequestCharacter

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
			characters = append(characters, element)
		}
		urlPage = request.Next

	}
	return characters
}
