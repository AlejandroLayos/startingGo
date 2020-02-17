package model

/*RequestCharacter result for the api get*/
type RequestCharacter struct {
	Count    int
	Next     string
	Previous string
	Results  []Character
}

/*RequestPlanet result for the api get*/
type RequestPlanet struct {
	Count    int
	Next     string
	Previous string
	Results  []Planet
}

/*RequestFilm result for the api get*/
type RequestFilm struct {
	Count    int
	Next     string
	Previous string
	Results  []Film
}

/*RequestSpecies result for the api get*/
type RequestSpecies struct {
	Count    int
	Next     string
	Previous string
	Results  []Species
}

/*RequestStarship result for the api get*/
type RequestStarship struct {
	Count    int
	Next     string
	Previous string
	Results  []Starships
}

/*RequestVehicle result for the api get*/
type RequestVehicle struct {
	Count    int
	Next     string
	Previous string
	Results  []Vehicle
}
