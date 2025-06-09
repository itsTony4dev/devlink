package routes

import (
	"devlink/internal/handlers"
	"devlink/internal/middleware"
	"github.com/gorilla/mux"
)

func RegisterUserRoutes(router *mux.Router, userHandler *handlers.UserHandler, authHandler *handlers.AuthHandler) {
	userRouter := router.PathPrefix("/users").Subrouter().StrictSlash(true)

	// Auth-related routes
	userRouter.HandleFunc("/register", authHandler.RegisterUserHandler).Methods("POST")
	userRouter.HandleFunc("/login", authHandler.LoginUserHandler).Methods("POST")
	userRouter.HandleFunc("/logout", authHandler.LogoutUserHandler).Methods("POST")
	
	// Protected routes for authenticated users
	protected := userRouter.NewRoute().Subrouter()
	// User-related routes
	protected.Use(middleware.JWTAuthMiddleware)
	protected.HandleFunc("/", userHandler.GetAllUsersHandler).Methods("GET")
	protected.HandleFunc("/{id}", userHandler.GetUserByIDHandler).Methods("GET")
	protected.HandleFunc("/{id}", userHandler.UpdateUserHandler).Methods("PUT")
	protected.HandleFunc("/{id}", userHandler.DeleteUserHandler).Methods("DELETE")
}
