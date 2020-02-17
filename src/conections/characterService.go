package conections

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/AlejandroLayos/startingGo/src/model"
)

var urlCharacter string = "https://swapi.co/api/people/"

//GetListOfCharacters gets all the Characters registered
func GetListOfCharacters() []model.Character {
	var request model.RequestCharacter
	var characters []model.Character
	spaceClient := http.Client{
		Timeout: time.Second * 5, // Maximum of 10 secs
	}

	req, err := http.NewRequest(http.MethodGet, urlCharacter, nil)
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
	characters = request.Results
	return characters
}

//GetCharacterByURL gets one Character depending of the url
func GetCharacterByURL(url string) model.Character {
	var character model.Character

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
	jsonErr := json.Unmarshal(body, &character)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return character
}
