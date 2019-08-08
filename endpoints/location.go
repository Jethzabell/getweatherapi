package endpoints

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"weather-app/models"

	"github.com/labstack/echo"
)

//GetLocation Main function this will get the zipcode from the url and return the forecast
func GetLocation(c echo.Context) error {

	zipCode := c.QueryParam("zipcode")

	// ------------ REQUEST ------------
	response, err := http.Get("https://api.mapbox.com/geocoding/v5/mapbox.places/" + zipCode + ".json?access_token=pk.eyJ1IjoiamV0aHphYmVsbCIsImEiOiJjanoyNTBmaDUwMm9kM25vYWhua2QwM3huIn0.jTENbgBTU-vlbDRIoXLSAA&limit=1")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}
	// ------------ END REQUEST ------------

	// ------------ UNMARSHALL ------------
	s, err1 := GetMapBoxResponse([]byte(body))
	if err1 != nil {
		panic(err1.Error())
	}

	latitude := s.Features[0].Center[0]
	longitude := s.Features[0].Center[1]
	// ------------ END UNMARSHALL ------------

	// ------------ GET FORECAST ------------
	forecast, erro := GetForeCast(latitude, longitude)
	if erro != nil {
		panic(erro.Error())
	}
	// ------------ END GET FORECAST ------------

	// ------------ CREATING RESPONSE ------------
	theResponse := models.Response{
		PlaceName:   s.Features[0].PlaceName,
		Temperature: forecast.Currently.Temperature,
		Summary:     forecast.Daily.Data[0].Summary,
	}
	// ------------ END CREATING RESPONSE ------------

	return c.JSON(http.StatusOK, theResponse)
}

//GetMapBoxResponse Unmarshal
func GetMapBoxResponse(body []byte) (*models.MapBoxResponse, error) {
	var s = new(models.MapBoxResponse)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}
