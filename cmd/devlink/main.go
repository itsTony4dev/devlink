package main

import (
	"log"
	"net/http"

	"devlink/internal/config"
	"devlink/internal/db"
	"devlink/internal/routes"
)

func main() {
	config.LoadEnv() 
	port := config.GetEnv("PORT", "8080")

	db.InitDB()

	r := routes.SetupRouter()
	log.Printf("Server is running on port %s", port)
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	log.Println("Server stopped")
}
