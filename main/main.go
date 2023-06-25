package main

import (
	"aytodo/database"
	"aytodo/server"
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {

	// Initialize the database
	if err := database.ConnectAndMigrate(
		"localhost",
		"5432",
		"todo",
		"postgres",
		"1337",
		database.SSLModeDisable); err != nil {
		logrus.Panicf("Failed to initialize and migrate database with error: %+v", err)
	}

	// Setup the routes
	server.SetupRoutes()
	// Start the server
	logrus.Print("Server started on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logrus.Panicf("Failed to start server with error: %+v", err)
	}

}
