package handlers

import "devlink/internal/repository"

type HandlersContainer struct {
	UserHandler     *UserHandler
	AuthHandler     *AuthHandler
	ResourceHandler *ResourceHandler
}

func NewHandlersContainer(userRepository *repository.UserRepository, resourceRepository *repository.ResourceRepository) *HandlersContainer {
	return &HandlersContainer{
		UserHandler:     NewUserHandler(userRepository),
		AuthHandler:     NewAuthHandler(userRepository),
		ResourceHandler: NewResourceHandler(resourceRepository),
	}
}
