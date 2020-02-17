package test

import (
	"testing"

	"github.com/AlejandroLayos/startingGo/src/service"
)

func TestGetCharacter(t *testing.T) {
	character := service.GetCharacter("Luke Skywalker")
	if character.Name == "" {
		t.Errorf("The Character is empty")
	}
}

func TestGetFilm(t *testing.T) {
	film := service.GetFilm(1)
	if film.Title == "" {
		t.Errorf("The film is empty")
	}
}
func TestGetPlanet(t *testing.T) {
	planet := service.GetPlanet("Alderaan")
	if planet.Name == "" {
		t.Errorf("The planet is empty")
	}
}
func TestGetSpecies(t *testing.T) {
	species := service.GetSpecies("Hutt")
	if species.Name == "" {
		t.Errorf("The species is empty")
	}
}
func TestGetStarship(t *testing.T) {
	starship := service.GetStarship("X-wing")
	if starship.Name == "" {
		t.Errorf("The starship is empty")
	}
}
func TestGetVehicle(t *testing.T) {
	vechicle := service.GetVehicle("Snowspeeder")
	if vechicle.Name == "" {
		t.Errorf("The vechicle is empty")
	}
}
