package database

import (
	"MyGram-Golang-DTS/helper"
	"MyGram-Golang-DTS/model"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	host := helper.ConfigValue("DB_HOST")
	port := helper.ConfigValue("DB_PORT")
	username := helper.ConfigValue("DB_USERNAME")
	password := helper.ConfigValue("DB_PASSWORD")
	dbName := helper.ConfigValue("DB_NAME")

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
