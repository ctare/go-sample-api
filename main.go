package main

import (
	"api/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
	"net/http"
)

type Node struct {
	gorm.Model
	Tag string `json:"tag"`
}

type Pair struct {
	gorm.Model

	// Node from to
	From uint `json:"from_"`
	To   uint `json:"to_"`
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db := database()

	err := migrate(db)
	if err != nil {
		e.Logger.Fatal(err)
	}

	controller.FruitInitialize(e, database)

	e.GET("/nodes", nodes)
	e.GET("/pairs", pairs)
	e.POST("/createPair", createPair)

	e.Logger.Fatal(e.Start(":8080"))
}

func nodes(c echo.Context) error {
	var nodes []Node
	if err := database().Find(&nodes).Error; err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, nodes)
}

func pairs(c echo.Context) error {
	var pairs []Pair
	if err := database().Find(&pairs).Error; err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, pairs)
}

func createPair(c echo.Context) error {
	// body as Pair json
	pair := new(Pair)
	if err := c.Bind(pair); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	// --- validation
	if pair.From == 0 || pair.To == 0 {
		return c.NoContent(http.StatusBadRequest)
	}

	// --- create
	database().Create(pair)
	return c.NoContent(http.StatusOK)
}
