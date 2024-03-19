package database

import (
	"MyGram-Golang-DTS/model"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file", err)
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		host, username, password, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db, err
}

func MigrateDB(db *gorm.DB) error {
	err := db.AutoMigrate(
		model.User{},
		model.SocialMedia{},
		model.Photo{},
		model.Comment{},
	)

	if err != nil {
		return err
	}
	return nil
}
