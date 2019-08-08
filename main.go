// By Jethzabell Medina
// GOLANG REST API Get ForeCast

package main

import (
	"fmt"
	"weather-app/endpoints"

	"github.com/labstack/echo"
)

var e *echo.Echo

func main() {
	fmt.Println("Welcome to the server")
	e := echo.New()
	e.GET("/", endpoints.GetLocation)
	e.Start(":1323")
}
