package model

/*Vehicle model*/
type Vehicle struct {
	Name                   string
	Model                  string
	Manufacturer           string
	Cost_in_credits        string
	Length                 string
	Max_atmosphering_speed string
	Crew                   string
	Passengers             string
	Cargo_capacity         string
	Consumables            string
	Vehicle_class          string
	Pilots                 []string
	Films                  []string
	Created                string
	Edited                 string
	Url                    string
}
