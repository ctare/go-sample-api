package controller

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type Fruit struct {
	gorm.Model
	Name string `json:"name"`
	Icon string `json:"icon"`
}

const base = "/fruits"

var database func() *gorm.DB

func FruitInitialize(e *echo.Echo, openDatabase func() *gorm.DB) {
	database = openDatabase
	routing(e)
}

func routing(e *echo.Echo) {
	e.GET(base, fruits)
	e.POST(base+"/create", create)
}

func fruits(c echo.Context) error {
	var fruits []Fruit
	database().Find(&fruits)
	return c.JSON(http.StatusOK, fruits)
}

func create(c echo.Context) error {
	// body as Fruit json
	fruit := new(Fruit)
	if err := c.Bind(fruit); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	// --- validation
	if fruit.Name == "" {
		return c.NoContent(http.StatusBadRequest)
	}

	// --- create
	database().Create(fruit)
	return c.NoContent(http.StatusOK)
}
