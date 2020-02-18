package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/AlejandroLayos/startingGo/src/service"

	"github.com/manifoldco/promptui"
)

func main() {
	fmt.Printf("||           Starting Star Wars Catalog         || ")
	fmt.Printf("Fetching data")
	service.DownloadData()
	fmt.Printf("Data retrieved succesfully!")
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
			fmt.Printf(string(prettyJSON))

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
			fmt.Printf(string(prettyJSON))

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
			fmt.Printf(string(prettyJSON))

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
			fmt.Printf(string(prettyJSON))

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
			fmt.Printf(string(prettyJSON))

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
			fmt.Printf(string(prettyJSON))
		case "-----Exit-----":
			fmt.Printf("||   May the Force be with you   ||")
			exist = true
		default:
			fmt.Printf("Too far away.")
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
				fmt.Printf("||   May the Force be with you   ||")
				return
			}

			exist = result2 != "y"

		}

	}

}
