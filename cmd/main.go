package main

import (
	"catalyst-players/internal/domain/entities"
	"catalyst-players/internal/infrastructure/database"
	"catalyst-players/internal/presentation/routes"
	"log"
	"net/http"
	"os"
)

func main() {
	// Initialize database connection
	dbConfig := database.NewConfig()
	db, err := database.Connect(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto migrate database tables
	err = db.AutoMigrate(
		&entities.Tag{},
		&entities.Stadium{},
		&entities.Team{},
		&entities.Player{},
		&entities.League{},
		&entities.Season{},
		&entities.Match{},
		&entities.MatchPlayer{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Setup routes
	router := routes.SetupRoutes(db)

	// Get server port from environment
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080" // Default port
	}

	// Start server
	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
