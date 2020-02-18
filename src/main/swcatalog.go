package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/AlejandroLayos/startingGo/src/service"

	"github.com/manifoldco/promptui"
)

func main() {
	fmt.Println("||           Starting Star Wars Catalog         || ")
	fmt.Println("Fetching data")
	service.DownloadData()
	fmt.Println("Data retrieved succesfully!")
	time.Sleep(1 * time.Second)
	var characters []string = service.GetListCharacter()
	var films []string = service.GetListFilms()
	var planets []string = service.GetListPlanets()
	var species []string = service.GetListSpecies()
	var starships []string = service.GetListStarships()
	var vehicles []string = service.GetListVechicles()
	var exist = false

	for exist == false {

		prompt := promptui.Select{
			Label: "What are you searching for?",
			Items: []string{"Planet", "Character", "Film", "Vehicle", "Starship",
				"Species", "-----Exit-----"},
		}

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		switch result {

		case "Planet":

			prompt := promptui.Select{

				Label: "Select one planet",
				Items: planets,
			}

			_, result, err := prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			planet := service.FindPlanetByName(result)
			prettyJSON, _ := json.MarshalIndent(planet, "", "\t")
			fmt.Println(string(prettyJSON))

		case "Character":

			prompt := promptui.Select{

				Label: "Select one character",
				Items: characters,
			}

			_, result, err := prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			character := service.FindCharacterByName(result)
			prettyJSON, _ := json.MarshalIndent(character, "", "\t")
			fmt.Println(string(prettyJSON))

		case "Film":
			prompt := promptui.Select{

				Label: "Select one film",
				Items: films,
			}

			_, result, err := prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			film := service.FindFilmBytitle(result)
			prettyJSON, _ := json.MarshalIndent(film, "", "\t")
			fmt.Println(string(prettyJSON))

		case "Species":
			prompt := promptui.Select{

				Label: "Select one species",
				Items: species,
			}

			_, result, err := prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			specie := service.FindSpeciesByName(result)
			prettyJSON, _ := json.MarshalIndent(specie, "", "\t")
			fmt.Println(string(prettyJSON))

		case "Starship":

			prompt := promptui.Select{

				Label: "Select one starship",
				Items: starships,
			}

			_, result, err := prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			starship := service.FindStarshipByName(result)
			prettyJSON, _ := json.MarshalIndent(starship, "", "\t")
			fmt.Println(string(prettyJSON))

		case "Vehicle":

			prompt := promptui.Select{

				Label: "Select one vehicle",
				Items: vehicles,
			}

			_, result, err := prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			vehicle := service.FindVehicleByName(result)
			prettyJSON, _ := json.MarshalIndent(vehicle, "", "\t")
			fmt.Println(string(prettyJSON))
		case "-----Exit-----":
			fmt.Println("||   May the Force be with you   ||")
			exist = true
		default:
			fmt.Println("Too far away.")
			exist = true
		}
		time.Sleep(2 * time.Second)
		if exist != true {

			prompt2 := promptui.Prompt{
				Label:     "Do you want to search anymore?",
				IsConfirm: true,
			}

			result2, err2 := prompt2.Run()

			if err2 != nil {
				fmt.Println("||   May the Force be with you   ||")
				return
			}

			exist = result2 != "y"

		}

	}

}

/*import (
	"encoding/json"
	"flag"
	"fmt"
	"strconv"

	"github.com/AlejandroLayos/startingGo/src/service"


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
*/
