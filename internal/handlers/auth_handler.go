package handlers

import (
	"devlink/internal/dto"
	"devlink/internal/models"
	"devlink/internal/repository"
	"devlink/internal/utils"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	repo *repository.UserRepository
}

func NewAuthHandler(userRepository *repository.UserRepository) *AuthHandler {
	return &AuthHandler{
		repo: userRepository,
	}
}

func (h *AuthHandler) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var registerReq dto.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&registerReq); err != nil {
		dto.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Create user model from request
	user := models.User{
		Username: registerReq.Username,
		Email:    registerReq.Email,
		Password: registerReq.Password,
	}

	// Validate user input
	if err := user.ValidateUsername(); err != nil {
		dto.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if err := user.ValidateEmail(); err != nil {
		dto.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if err := user.ValidatePassword(); err != nil {
		dto.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Check if user already exists
	if existingUser, _ := h.repo.GetByEmail(user.Email); existingUser != nil {
		dto.WriteError(w, http.StatusConflict, models.ErrEmailExists)
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		dto.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	user.Password = string(hashedPassword)

	// Create user
	if err := h.repo.CreateUser(&user); err != nil {
		dto.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID, user.Email, user.Username)
	if err != nil {
		dto.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	dto.WriteSuccess(w, http.StatusCreated, map[string]interface{}{
		"user":  user.ToResponse(),
		"token": token,
	}, "User registered successfully")
}

func (h *AuthHandler) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	var loginReq dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		dto.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Validate email format
	user := models.User{Email: loginReq.Email}
	if err := user.ValidateEmail(); err != nil {
		dto.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Get user from database
	user, err := h.repo.GetByEmail(loginReq.Email)
	if err != nil {
		dto.WriteError(w, http.StatusUnauthorized, models.ErrInvalidCredentials)
		return
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password)); err != nil {
		dto.WriteError(w, http.StatusUnauthorized, models.ErrInvalidCredentials)
		return
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID, user.Email, user.Username)
	if err != nil {
		dto.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	dto.WriteSuccess(w, http.StatusOK, map[string]interface{}{
		"user":  user.ToResponse(),
		"token": token,
	}, "Login successful")
}

func (h *AuthHandler) LogoutUserHandler(w http.ResponseWriter, r *http.Request) {
	dto.WriteSuccess(w, http.StatusOK, nil, "User logged out successfully")
}
