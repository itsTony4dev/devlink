package main

import (
	"log"
	"net/http"

	"devlink/internal/config"
	"devlink/internal/db"
	"devlink/internal/handlers"
	"devlink/internal/repository"
	"devlink/internal/routes"
)

func main() {
	config.LoadEnv()
	port := config.GetEnv("PORT", "8080")

	dbConn := db.InitDB()

	userRepo := repository.NewUserRepository(dbConn)

	handlers := handlers.NewHandlersContainer(userRepo)

	r := routes.SetupRouter(handlers)

	log.Printf("Server is running on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	log.Println("Server stopped")
}
