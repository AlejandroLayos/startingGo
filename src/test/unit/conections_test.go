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
