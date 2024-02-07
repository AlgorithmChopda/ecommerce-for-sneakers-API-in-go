package main

import (
	"context"
	"fmt"

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

	fmt.Println("Database connected successfully ", conn)
}
