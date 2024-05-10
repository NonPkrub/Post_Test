package main

import (
	"Test/api/route"
	"Test/database"
	"Test/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var err error

func main() {
	database.DB, err = gorm.Open(postgres.Open(database.DbConnect(database.BuildDBConfig())), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	app := route.SetupRouter()
	database.DB.AutoMigrate(&domain.Post{})
	err = app.Listen((":8000"))
	if err != nil {
		panic(err)
	}
}
