package models

type Response struct {
	PlaceName   string  `json:"place_name"`
	Temperature float64 `json:"temperature"`
	Summary     string  `json:"summary"`
}
