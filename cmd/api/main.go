package main

import (
	"main/internal/server"

	"github.com/joho/godotenv"
)

// startServer is a helper function that starts the echo server.
func startServer() {
	server := server.NewServer()

	// Start echo server on port 3000
	server.Logger.Fatal(
		server.Start(":3000"))
}

// main is the entry point for the application.
func main() {
	godotenv.Load()

	startServer()
}
