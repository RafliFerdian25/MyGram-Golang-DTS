package main

import (
	"MyGram-Golang-DTS/database"
	routes "MyGram-Golang-DTS/router"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
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

	app := routes.New(db)

	err = godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file", err)
		log.Fatal("Error loading .env file")
	}

	apiPort := os.Getenv("API_PORT")
	log.Fatal(app.Run(apiPort))
}
