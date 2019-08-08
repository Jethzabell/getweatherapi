package models

type MapBoxResponse struct {
	Features []Features `json:"features"`
}

type Features struct {
	Center    []float64 `json:"center"`
	PlaceName string    `json:"place_name"`
}

type Center struct {
}
