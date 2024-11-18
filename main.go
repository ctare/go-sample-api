package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

type Fruit struct {
	gorm.Model
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type Node struct {
	gorm.Model
}

type Pair struct {
	gorm.Model

	// Node from to
	From uint `json:"from_"`
	To   uint `json:"to_"`
}

var Db *gorm.DB

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/fruits", fruits)
	e.POST("/create", create)
	e.GET("/nodes", nodes)
	e.POST("/createPair", createPair)

	dsn := "root:root@tcp(db:3306)/main?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		Db = db
	}

	err = db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Fruit{})
	if err != nil {
		return
	}

	err = db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Node{})
	if err != nil {
		return
	}

	err = db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Pair{})
	if err != nil {
		return
	}

	//db.Create(&Fruit{Name: "apple", Icon: "🍎"})

	e.Logger.Fatal(e.Start(":8080"))
}

func fruits(c echo.Context) error {
	var fruits []Fruit
	Db.Find(&fruits)
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
	Db.Create(fruit)
	return c.NoContent(http.StatusOK)
}

func nodes(c echo.Context) error {
	var nodes []Node
	Db.Find(&nodes)
	return c.JSON(http.StatusOK, nodes)
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
	Db.Create(pair)
	return c.NoContent(http.StatusOK)
}
