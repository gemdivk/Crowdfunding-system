package main

import (
	"github.com/gemdivk/Crowdfunding-system/internal/db"
	"github.com/gemdivk/Crowdfunding-system/internal/routes"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db.InitDB()
	db.CheckDBConnection()
	router := routes.SetupRouter()

	log.Println("Starting server on port 8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
