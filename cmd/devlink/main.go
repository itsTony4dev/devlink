package main

import (
	"log"
	"net/http"

	"github.com/itsTony4dev/devlink/internal/config"
	"github.com/itsTony4dev/devlink/internal/routes"
)

func main() {
	config.LoadEnv() // Load environment variables
	port := config.GetEnv("PORT", "8080")

	r := routes.SetupRouter()
	log.Printf("Server is running on port %s", port)
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	log.Println("Server stopped")
}
