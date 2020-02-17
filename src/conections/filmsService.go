package conections

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/AlejandroLayos/startingGo/src/model"
)

var urlFilms string = "https://swapi.co/api/films/"

//GetListOfFilms gets all the Films registered
func GetListOfFilms() []model.Film {
	var request model.RequestFilm
	var films []model.Film
	spaceClient := http.Client{
		Timeout: time.Second * 5, // Maximum of 10 secs
	}

	req, err := http.NewRequest(http.MethodGet, urlFilms, nil)
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
	films = request.Results
	return films
}

//GetFilmByURL gets one Film depending of the url
func GetFilmByURL(url string) model.Film {
	var film model.Film
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
	jsonErr := json.Unmarshal(body, &film)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return film
}
