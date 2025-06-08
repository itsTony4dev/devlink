package routes

import (
	"devlink/internal/handlers"
	"github.com/gorilla/mux"
)

func RegisterUserRoutes(router *mux.Router, userHandler *handlers.UserHandler, authHandler *handlers.AuthHandler) {
	userRouter := router.PathPrefix("/users").Subrouter().StrictSlash(true)

	// Auth-related routes
	userRouter.HandleFunc("/register", authHandler.RegisterUserHandler).Methods("POST")
	userRouter.HandleFunc("/login", authHandler.LoginUserHandler).Methods("POST")
	userRouter.HandleFunc("/logout", authHandler.LogoutUserHandler).Methods("POST")

	// User-related routes
	userRouter.HandleFunc("/", userHandler.GetAllUsersHandler).Methods("GET")
	userRouter.HandleFunc("/{id}", userHandler.GetUserByIDHandler).Methods("GET")
	userRouter.HandleFunc("/{id}", userHandler.UpdateUserHandler).Methods("PUT")
	userRouter.HandleFunc("/{id}", userHandler.DeleteUserHandler).Methods("DELETE")
}
