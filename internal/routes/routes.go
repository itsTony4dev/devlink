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
		[]string{"http://localhost:3000"}, // Add your frontend origins
		[]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		[]string{"Content-Type", "Authorization"},
	)

	// Apply global middlewares
	r.Use(middleware.SecurityHeaders)
	r.Use(rateLimiter.RateLimit)
	r.Use(corsMiddleware.CORS)

	// Health check endpoint
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("DevLink is running 🚀!"))
	}).Methods("GET")

	// Register user routes
	RegisterUserRoutes(r, h.UserHandler, h.AuthHandler)

	return r
}
