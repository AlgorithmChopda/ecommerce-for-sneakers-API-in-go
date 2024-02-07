package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/api"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/app"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/repository"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Server starting...")
	defer fmt.Println("Server stopped")

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading env file")
		return
	}

	ctx := context.TODO()
	conn, err := repository.InitializeDB(ctx)
	if err != nil {
		fmt.Println("Error connecting to Database ", err)
		return
	}

	fmt.Println("Database connected successfully")

	// Initialize Dependencies
	deps := app.NewService(conn)
	router := api.NewRouter(deps)

	err = http.ListenAndServe("127.0.0.1:8000", router)
	if err != nil {
		fmt.Println(err)
		return
	}
}
