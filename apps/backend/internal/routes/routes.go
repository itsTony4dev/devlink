package routes

import (
	"devlink/internal/handlers"
	"devlink/internal/middleware"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func SetupRouter(h *handlers.HandlersContainer) *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	// Create middleware instances
	rateLimiter := middleware.NewRateLimiter(100, time.Minute) // 100 requests per minute
	corsMiddleware := middleware.NewCORSMiddleware(
		[]string{"http://localhost:5173", "http://127.0.0.1:5173"}, // Explicitly allowed frontend origins
		[]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		[]string{"Content-Type", "Authorization"},
	)

	// Apply global middlewares
	r.Use(middleware.SecurityHeaders)
	r.Use(rateLimiter.RateLimit)
	r.Use(corsMiddleware.CORS)

	// Global OPTIONS handler for preflight requests
	// This will be handled by the CORS middleware, but kept here for clarity if needed elsewhere.
	r.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Health check endpoint
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("DevLink is running 🚀!"))
	}).Methods("GET")

	// Register user routes
	RegisterUserRoutes(r, h.UserHandler, h.AuthHandler)

	// Register resource routes
	RegisterResourceRoutes(r, h.ResourceHandler)

	return r
}
