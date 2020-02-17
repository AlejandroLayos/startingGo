package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"strconv"

	"github.com/AlejandroLayos/startingGo/src/service"
)

func main() {
	//This flag is used to detect the type of the object to search
	typePtr := flag.String("type", "", "type --> planet, character, species, starship, vehicles (string), episode (number)")
	flag.Parse()

	var nameToFind string
	//For used to build the name from the array
	for index, element := range flag.Args() {
		if index == 0 {
			nameToFind = element
		} else {
			nameToFind = nameToFind + " " + element
		}
	}
	switch *typePtr {
	case "planet":
		planet := service.GetPlanet(nameToFind)
		if planet.Name == "" {
			fmt.Println("The planet does not exist in our database")
		} else {
			prettyJSON, _ := json.MarshalIndent(planet, "", "\t")
			fmt.Println(string(prettyJSON))
		}
	case "character":
		character := service.GetCharacter(nameToFind)
		if character.Name == "" {
			fmt.Println("The character does not exist in our database")

		} else {
			prettyJSON, _ := json.MarshalIndent(character, "", "\t")
			fmt.Println(string(prettyJSON))
		}
	case "episode":
		episode, _ := strconv.Atoi(nameToFind)
		film := service.GetFilm(episode)
		if film.Title == "" {
			fmt.Println("The film does not exist in our database")
		} else {
			prettyJSON, _ := json.MarshalIndent(film, "", "\t")
			fmt.Println(string(prettyJSON))
		}
	case "species":
		species := service.GetSpecies(nameToFind)
		if species.Name == "" {
			fmt.Println("The species does not exist in our database")
		} else {
			prettyJSON, _ := json.MarshalIndent(species, "", "\t")
			fmt.Println(string(prettyJSON))
		}
	case "starship":
		starship := service.GetStarship(nameToFind)
		if starship.Name == "" {
			fmt.Println("The starship does not exist in our database")
		} else {
			prettyJSON, _ := json.MarshalIndent(starship, "", "\t")
			fmt.Println(string(prettyJSON))
		}
	case "vehicles":
		vehicle := service.GetVehicle(nameToFind)
		if vehicle.Name == "" {
			fmt.Println("The vechicle does not exist in our database")
		} else {
			prettyJSON, _ := json.MarshalIndent(vehicle, "", "\t")
			fmt.Println(string(prettyJSON))
		}

	default:
		fmt.Println("Too far away.")
	}

}
