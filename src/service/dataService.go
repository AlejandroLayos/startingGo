package service

import (
	"fmt"
	"strings"

	"github.com/AlejandroLayos/startingGo/src/conections"
	"github.com/AlejandroLayos/startingGo/src/model"
)

var characterList []model.Character
var planetList []model.Planet
var filmList []model.Film
var vehicleList []model.Vehicle
var starshipList []model.Starships
var speciesList []model.Species

//DownloadData gets all the data form de app
func DownloadData() {
	fmt.Println("characters...")
	characterList = conections.GetListOfCharacters()

	fmt.Println("planets..")
	planetList = conections.GetListOfPlanets()

	fmt.Println("films...")
	filmList = conections.GetListOfFilms()
	fmt.Println("vehicles...")
	vehicleList = conections.GetListOfVehicles()

	fmt.Println("starships...")
	starshipList = conections.GetListOfStarships()

	fmt.Println("species...")
	speciesList = conections.GetListOfSpecies()

	updatePlanets()
	updateFilms()
	updateStarships()
	updateVehicles()
	updateSpecies()
	updateCharacters()
}

//GetListCharacter return the string list of the characters
func GetListCharacter() []string {
	var nameList []string
	for _, element := range characterList {
		nameList = append(nameList, element.Name)
	}
	return nameList
}

//GetListFilms return the string list of the films
func GetListFilms() []string {
	var titleList []string
	for _, element := range filmList {
		titleList = append(titleList, element.Title)
	}
	return titleList
}

//GetListPlanets return the string list of the planets
func GetListPlanets() []string {
	var nameList []string
	for _, element := range planetList {
		nameList = append(nameList, element.Name)
	}
	return nameList
}

//GetListVechicles return the string list of the vehicles
func GetListVechicles() []string {
	var nameList []string
	for _, element := range vehicleList {
		nameList = append(nameList, element.Name)
	}
	return nameList
}

//GetListStarships return the string list of the starships
func GetListStarships() []string {
	var nameList []string
	for _, element := range starshipList {
		nameList = append(nameList, element.Name)
	}
	return nameList
}

//GetListSpecies return the string list of the species
func GetListSpecies() []string {
	var nameList []string
	for _, element := range speciesList {
		nameList = append(nameList, element.Name)
	}
	return nameList
}

//FindCharacterByName returns a character by name
func updateCharacters() {

	for _, element := range characterList {
		character := element

		if character.Name != "" {
			for index, element := range character.Films {
				if strings.Contains(element, "http") {
					character.Films[index] = findFilmByURL(element)
				}
			}
			for index, element := range character.Species {
				if strings.Contains(element, "http") {
					character.Species[index] = findSpeciesByURL(element)
				}
			}
			for index, element := range character.Vehicles {
				if strings.Contains(element, "http") {
					character.Vehicles[index] = findVehicleByURL(element)
				}
			}
			for index, element := range character.Starships {
				if strings.Contains(element, "http") {
					character.Starships[index] = findStarshipByURL(element)
				}
			}
		}
	}
}

//FindCharacterByName returns a character by name
func FindCharacterByName(name string) model.Character {
	var character model.Character
	for _, element := range characterList {
		if element.Name == name {
			character = element
			break
		}
	}
	if strings.Contains(character.Homeworld, "http") {
		character.Homeworld = findPlanetByURL(character.Homeworld)

	}
	return character
}

func findCharacterByURL(url string) string {
	var character model.Character
	for _, element := range characterList {
		if element.Url == url {
			character = element
			break
		}
	}
	return character.Name
}

func updatePlanets() {
	for _, element := range planetList {
		planet := element
		if planet.Name != "" {
			for index, element := range planet.Residents {
				if strings.Contains(element, "http") {
					planet.Residents[index] = findCharacterByURL(element)
				}
			}
			for index, element := range planet.Films {
				if strings.Contains(element, "http") {
					planet.Films[index] = findFilmByURL(element)
				}
			}
		}

	}

}

//FindPlanetByName returns a Planet by name
func FindPlanetByName(name string) model.Planet {
	var planet model.Planet
	for _, element := range planetList {
		if element.Name == name {
			planet = element
			break
		}
	}
	return planet
}

func findPlanetByURL(url string) string {
	var planet model.Planet
	for _, element := range planetList {
		if element.Url == url {
			planet = element
			break
		}
	}

	return planet.Name
}

func updateFilms() {
	for _, element := range filmList {
		film := element
		if film.Title != "" {
			for index, element := range film.Characters {
				film.Characters[index] = findCharacterByURL(element)
			}
			for index, element := range film.Planets {
				if strings.Contains(element, "http") {
					film.Planets[index] = findPlanetByURL(element)
				}
			}
			for index, element := range film.Starships {
				if strings.Contains(element, "http") {
					film.Starships[index] = findStarshipByURL(element)
				}
			}
			for index, element := range film.Vehicles {
				if strings.Contains(element, "http") {
					film.Vehicles[index] = findVehicleByURL(element)
				}
			}
			for index, element := range film.Species {
				if strings.Contains(element, "http") {
					film.Species[index] = findSpeciesByURL(element)
				}
			}
		}

	}

}

//FindFilmBytitle returns a film by title
func FindFilmBytitle(title string) model.Film {
	var film model.Film
	for _, element := range filmList {
		if element.Title == title {
			film = element
			break
		}
	}
	return film
}
func findFilmByURL(url string) string {
	var film model.Film
	for _, element := range filmList {
		if element.Url == url {
			film = element
			break
		}
	}

	return film.Title
}

func updateVehicles() {
	for _, element := range vehicleList {
		vehicle := element
		if vehicle.Name != "" {
			for index, element := range vehicle.Pilots {
				if strings.Contains(element, "http") {
					vehicle.Pilots[index] = findCharacterByURL(element)
				}
			}
			for index, element := range vehicle.Films {
				if strings.Contains(element, "http") {
					vehicle.Films[index] = findFilmByURL(element)
				}
			}
		}
	}
}

//FindVehicleByName returns a vehicle by name
func FindVehicleByName(name string) model.Vehicle {
	var vehicle model.Vehicle
	for _, element := range vehicleList {
		if element.Name == name {
			vehicle = element
			break
		}
	}
	return vehicle
}
func findVehicleByURL(url string) string {
	var vehicle model.Vehicle
	for _, element := range vehicleList {
		if element.Url == url {
			vehicle = element
			break
		}
	}
	return vehicle.Name
}

//FindStarshipByName returns a starship by name
func FindStarshipByName(name string) model.Starships {
	var starship model.Starships

	for _, element := range starshipList {
		if element.Name == name {
			starship = element
			break
		}
	}
	if starship.Name != "" {
		for index, element := range starship.Pilots {
			if strings.Contains(element, "http") {
				starship.Pilots[index] = findCharacterByURL(element)
			}
		}

		for index, element := range starship.Films {
			if strings.Contains(element, "http") {
				starship.Films[index] = findFilmByURL(element)
			}
		}
	}

	return starship
}
func updateStarships() {
	for _, element := range starshipList {
		starship := element
		if starship.Name != "" {
			for index, element := range starship.Pilots {
				if strings.Contains(element, "http") {
					starship.Pilots[index] = findCharacterByURL(element)
				}
			}

			for index, element := range starship.Films {
				if strings.Contains(element, "http") {
					starship.Films[index] = findFilmByURL(element)
				}
			}
		}
	}
}

func findStarshipByURL(url string) string {
	var starship model.Starships
	for _, element := range starshipList {
		if element.Url == url {
			starship = element
			break
		}
	}
	return starship.Name
}
func updateSpecies() {
	for _, element := range speciesList {
		species := element
		if species.Name != "" {
			species.Homeworld = findPlanetByURL(species.Homeworld)
			for index, element := range species.People {
				if strings.Contains(element, "http") {
					species.People[index] = findCharacterByURL(element)
				}
			}

			for index, element := range species.Films {
				if strings.Contains(element, "http") {
					species.Films[index] = findFilmByURL(element)
				}
			}
		}
	}
}

//FindSpeciesByName returns a species by name
func FindSpeciesByName(name string) model.Species {
	var species model.Species
	for _, element := range speciesList {
		if element.Name == name {
			species = element
			break
		}
	}
	if species.Name != "" {
		species.Homeworld = findPlanetByURL(species.Homeworld)
		for index, element := range species.People {
			if strings.Contains(element, "http") {
				species.People[index] = findCharacterByURL(element)
			}
		}

		for index, element := range species.Films {
			if strings.Contains(element, "http") {
				species.Films[index] = findFilmByURL(element)
			}
		}
	}

	return species
}

func findSpeciesByURL(url string) string {
	var species model.Species
	for _, element := range speciesList {
		if element.Url == url {
			species = element
			break
		}
	}
	return species.Name
}
