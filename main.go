package main

import (
	"MyGram-Golang-DTS/database"
	"MyGram-Golang-DTS/helper"
	"MyGram-Golang-DTS/routes"
	"log"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("Error connect to database", err)
		panic(err)
	}

	err = database.MigrateDB(db)
	if err != nil {
		log.Fatal("Error migrate to database", err)
		panic(err)
	}

	app := routes.NewRoute(db)

	apiPort := helper.ConfigValue("PORT")
	log.Fatal(app.Run(":" + apiPort))
}
