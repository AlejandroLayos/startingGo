package test

import (
	"testing"

	"github.com/AlejandroLayos/startingGo/src/conections"
)

func TestGetCharacterList(t *testing.T) {
	characterList := conections.GetListOfCharacters()
	if len(characterList) < 1 {
		t.Errorf("The list is empty")
	}
}
func TestGetFilmList(t *testing.T) {
	characterList := conections.GetListOfFilms()
	if len(characterList) < 1 {
		t.Errorf("The list is empty")
	}
}

func TestGetPlanetList(t *testing.T) {
	characterList := conections.GetListOfPlanets()
	if len(characterList) < 1 {
		t.Errorf("The list is empty")
	}
}

func TestGetSpeciesList(t *testing.T) {
	characterList := conections.GetListOfSpecies()
	if len(characterList) < 1 {
		t.Errorf("The list is empty")
	}
}

func TestGetStarshipList(t *testing.T) {
	characterList := conections.GetListOfStarships()
	if len(characterList) < 1 {
		t.Errorf("The list is empty")
	}
}

func TestGetVechicleList(t *testing.T) {
	characterList := conections.GetListOfVehicles()
	if len(characterList) < 1 {
		t.Errorf("The list is empty")
	}
}

func TestGetCharacterByURL(t *testing.T) {
	character := conections.GetCharacterByURL("https://swapi.co/api/people/1/")
	if character.Name == "" {
		t.Errorf("The Character is empty")
	}
}
func TestGetFilmByURL(t *testing.T) {
	film := conections.GetFilmByURL("https://swapi.co/api/films/3/")
	if film.Title == "" {
		t.Errorf("The film is empty")
	}
}
func TestGetPlanetByURL(t *testing.T) {
	planet := conections.GetPlanetByURL("https://swapi.co/api/planets/2/")
	if planet.Name == "" {
		t.Errorf("The planet is empty")
	}
}
func TestGetSpeciesByURL(t *testing.T) {
	species := conections.GetSpeciesByURL("https://swapi.co/api/species/5/")
	if species.Name == "" {
		t.Errorf("The species is empty")
	}
}
func TestGetStarshipByURL(t *testing.T) {
	starship := conections.GetStarshipsByURL("https://swapi.co/api/starships/12/")
	if starship.Name == "" {
		t.Errorf("The starship is empty")
	}
}
func TestGetVehicleByURL(t *testing.T) {
	vechicle := conections.GetVehicleByURL("https://swapi.co/api/vehicles/14/")
	if vechicle.Name == "" {
		t.Errorf("The vechicle is empty")
	}
}
