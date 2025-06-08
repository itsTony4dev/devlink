package routes

import (
	"devlink/internal/handlers"
	"devlink/internal/middleware"
	"github.com/gorilla/mux"
)

func RegisterResourceRoutes(router *mux.Router, resourceHandler *handlers.ResourceHandler) {
	resourceRouter := router.PathPrefix("/resources").Subrouter().StrictSlash(true)

	// Protected routes for authenticated users
	resourceRouter.Use(middleware.JWTAuthMiddleware)

	// Resource CRUD routes
	resourceRouter.HandleFunc("", resourceHandler.CreateResourceHandler).Methods("POST")
	resourceRouter.HandleFunc("", resourceHandler.GetUserResourcesHandler).Methods("GET")
	resourceRouter.HandleFunc("/{id}", resourceHandler.GetResourceByIDHandler).Methods("GET")
	resourceRouter.HandleFunc("/{id}", resourceHandler.UpdateResourceHandler).Methods("PUT")
	resourceRouter.HandleFunc("/{id}", resourceHandler.DeleteResourceHandler).Methods("DELETE")

	// Search and filter routes
	resourceRouter.HandleFunc("/search", resourceHandler.SearchResourcesHandler).Methods("GET")
	resourceRouter.HandleFunc("/tags", resourceHandler.GetResourcesByTagsHandler).Methods("GET")
} 