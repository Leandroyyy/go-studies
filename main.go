package main

import (
	"exemplo/matematica"
	"fmt"

	"github.com/labstack/echo/v4"
)

type Car struct {
	Name  string
	Model string
}

var cars []Car

func generateCars() {
	cars = append(cars, Car{Name: "Ferrari", Model: "VW"})
	cars = append(cars, Car{Name: "Porsche", Model: "3"})
	cars = append(cars, Car{Name: "Honda fit", Model: "honda"})
}

func main() {
	generateCars()
	e := echo.New()
	e.GET("/cars", getCars)
	e.POST("/cars", createCar)
	e.Logger.Fatal(e.Start(":8080"))
	fmt.Println(matematica.Soma(1, 2))
}

func getCars(c echo.Context) error {

	return c.JSON(200, cars)
}

func createCar(c echo.Context) error {
	car := new(Car)

	if err := c.Bind(car); err != nil {
		return err
	}

	cars = append(cars, *car)
	return c.JSON(200, cars)
}
