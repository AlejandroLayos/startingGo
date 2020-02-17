package model

/*Film model*/
type Film struct {
	Title         string
	Episode_id    int
	Opening_crawl string
	Director      string
	Producer      string
	Release_date  string
	Characters    []string
	Planets       []string
	Starships     []string
	Vehicles      []string
	Species       []string
	Created       string
	Edited        string
	Url           string
}
