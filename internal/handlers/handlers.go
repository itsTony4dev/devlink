package handlers

import "devlink/internal/repository"

type HandlersContainer struct {
	UserHandler *UserHandler
	AuthHandler *AuthHandler
}

func NewHandlersContainer(userRepository *repository.UserRepository) *HandlersContainer {
	return &HandlersContainer{
		UserHandler: NewUserHandler(userRepository),
		AuthHandler: NewAuthHandler(userRepository),
	}
}
