package models

type DarkSkyResponse struct {
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Daily     Daily     `json:"daily"`
	Currently Currently `json:"currently"`
}

type Daily struct {
	Summary string `json:"summary"`
	Data    []Data `json:"data"`
}

type Data struct {
	Summary string `json:"summary"`
}

type Currently struct {
	Temperature float64 `json:"temperature"`
}
