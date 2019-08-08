package endpoints

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"weather-app/models"
)

//GetForeCast will return the forecast base on a latitude and longitude
func GetForeCast(longitude float64, latitude float64) (*models.DarkSkyResponse, error) {

	long := fmt.Sprintf("%f", longitude)
	lat := fmt.Sprintf("%f", latitude)

	response, err := http.Get("https://api.darksky.net/forecast/85136be32e2fc6cd3f6cefabddf65873/" + lat + "," + long)
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}

	s, err := GetDarkSkyResponse([]byte(body))

	return s, nil
}

// GetDarkSkyResponse Unmarshall
func GetDarkSkyResponse(body []byte) (*models.DarkSkyResponse, error) {
	var s = new(models.DarkSkyResponse)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}
