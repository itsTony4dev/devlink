package routes

import (
	"devlink/internal/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRouter(h *handlers.HandlersContainer) *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("DevLink is running 🚀!"))
	}).Methods("GET")

	// Register user routes
	RegisterUserRoutes(r, h.UserHandler, h.AuthHandler)

	return r
}
