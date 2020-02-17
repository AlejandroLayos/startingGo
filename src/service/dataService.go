package service

import (
	"github.com/AlejandroLayos/startingGo/src/conections"
	"github.com/AlejandroLayos/startingGo/src/model"
)

var characterList []model.Character
var planetList []model.Planet
var filmList []model.Film
var vehicleList []model.Vehicle
var starshipList []model.Starships
var speciesList []model.Species

// GetCharacter returns the data of the character request */
func GetCharacter(characterName string) model.Character {
	characterList = conections.GetListOfCharacters()
	var character = findCharacterByName(characterName)
	return character
}

// GetPlanet returns the data of the planet request */
func GetPlanet(planetName string) model.Planet {
	planetList = conections.GetListOfPlanets()
	var planet = findPlanetByName(planetName)
	return planet
}

// GetFilm returns the data of the film request */
func GetFilm(episode int) model.Film {
	filmList = conections.GetListOfFilms()
	var film = findFilmByEpisode(episode)
	return film
}

// GetVehicle returns the data of the vechicle request */
func GetVehicle(name string) model.Vehicle {
	vehicleList = conections.GetListOfVehicles()
	var vehicle = findVehicleByName(name)
	return vehicle
}

// GetStarship returns the data of the starship request */
func GetStarship(name string) model.Starships {
	starshipList = conections.GetListOfStarships()
	var starship = findStarshipByName(name)
	return starship
}

// GetSpecies returns the data of the species request */
func GetSpecies(name string) model.Species {
	speciesList = conections.GetListOfSpecies()
	var species = findSpeciesByName(name)
	return species
}

func findCharacterByName(name string) model.Character {
	var character model.Character
	for _, element := range characterList {
		if element.Name == name {
			character = element
			break
		}
	}
	if character.Name != "" {
		character.Homeworld = conections.GetPlanetByURL(character.Homeworld).Name
		for index, element := range character.Films {
			character.Films[index] = conections.GetFilmByURL(element).Title
		}
		for index, element := range character.Species {
			character.Species[index] = conections.GetSpeciesByURL(element).Name
		}
		for index, element := range character.Vehicles {
			character.Vehicles[index] = conections.GetVehicleByURL(element).Name
		}
		for index, element := range character.Starships {
			character.Starships[index] = conections.GetStarshipsByURL(element).Name
		}
	}

	return character
}

func findPlanetByName(name string) model.Planet {
	var planet model.Planet
	for _, element := range planetList {
		if element.Name == name {
			planet = element
			break
		}
	}
	if planet.Name != "" {
		for index, element := range planet.Residents {
			planet.Residents[index] = conections.GetCharacterByURL(element).Name
		}
		for index, element := range planet.Films {
			planet.Films[index] = conections.GetFilmByURL(element).Title
		}

	}

	return planet
}

func findFilmByEpisode(episode int) model.Film {
	var film model.Film
	for _, element := range filmList {
		if element.Episode_id == episode {
			film = element
			break
		}
	}
	if film.Title != "" {
		for index, element := range film.Characters {
			film.Characters[index] = conections.GetCharacterByURL(element).Name
		}
		for index, element := range film.Planets {
			film.Planets[index] = conections.GetPlanetByURL(element).Name
		}
		for index, element := range film.Starships {
			film.Starships[index] = conections.GetPlanetByURL(element).Name
		}
		for index, element := range film.Vehicles {
			film.Vehicles[index] = conections.GetPlanetByURL(element).Name
		}
		for index, element := range film.Species {
			film.Species[index] = conections.GetPlanetByURL(element).Name
		}
	}

	return film
}

func findVehicleByName(name string) model.Vehicle {
	var vehicle model.Vehicle

	for _, element := range vehicleList {
		if element.Name == name {
			vehicle = element
			break
		}
	}
	if vehicle.Name != "" {
		for index, element := range vehicle.Pilots {
			vehicle.Pilots[index] = conections.GetCharacterByURL(element).Name
		}

		for index, element := range vehicle.Films {
			vehicle.Films[index] = conections.GetFilmByURL(element).Title
		}
	}

	return vehicle
}

func findStarshipByName(name string) model.Starships {
	var starship model.Starships

	for _, element := range starshipList {
		if element.Name == name {
			starship = element
			break
		}
	}
	if starship.Name != "" {
		for index, element := range starship.Pilots {
			starship.Pilots[index] = conections.GetCharacterByURL(element).Name
		}

		for index, element := range starship.Films {
			starship.Films[index] = conections.GetFilmByURL(element).Title
		}
	}

	return starship
}
func findSpeciesByName(name string) model.Species {
	var species model.Species
	for _, element := range speciesList {
		if element.Name == name {
			species = element
			break
		}
	}
	if species.Name != "" {
		for index, element := range species.People {
			species.People[index] = conections.GetCharacterByURL(element).Name
		}

		for index, element := range species.Films {
			species.Films[index] = conections.GetFilmByURL(element).Title
		}
	}

	return species
}
