package main

import (
	"MyGram-Golang-DTS/database"
	"fmt"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to database")

	err = database.MigrateDB(db)
	if err != nil {
		panic(err)
	}

	fmt.Println("Migrate to database")

	// // Setup Repo
	// orderRepo := order.NewOrderRepository(db)

	// // Setup Service
	// orderService := orderService.NewService(orderRepo)

	// // Setup Handler
	// orderHandler := orderHandler.NewHandler(orderService)

	// handler.NewHttpServer(orderHandler)
}
