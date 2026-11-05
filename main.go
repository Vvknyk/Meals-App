package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"meals_backend/config"
	"meals_backend/routes"
)

func main() {

	// Load environment variables
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Connect to PostgreSQL
	config.ConnectDB()
	defer config.DB.Close()

	// Create HTTP server
	mux := http.NewServeMux()

	// Register routes
	routes.RegisterRecipeRoutes(mux)

	// Start server
	log.Println("Meals API running on http://localhost:8080")

	err = http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Fatal("Server error:", err)
	}
}
