package main

import (
	"api/controller"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DSN = "root:root@tcp(db:3306)/main?charset=utf8mb4&parseTime=True"

func database() *gorm.DB {
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&controller.Fruit{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&Node{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&Pair{})
	if err != nil {
		return err
	}

	return nil

}
