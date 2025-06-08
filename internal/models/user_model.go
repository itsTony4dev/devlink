package models

import (
	"devlink/internal/dto"
	"regexp"
	"unicode"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string     `json:"username" gorm:"not null;uniqueIndex" validate:"required,min=3,max=50,alphanum"`
	Email     string     `json:"email" gorm:"not null;uniqueIndex" validate:"required,email"`
	Password  string     `json:"password" gorm:"not null" validate:"required,min=8"`
	Resources []Resource `json:"resources"`
}

// ValidatePassword checks if the password meets the requirements
func (u *User) ValidatePassword() error {
	var (
		hasMinLen  = len(u.Password) >= 8
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	for _, char := range u.Password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasMinLen || !hasUpper || !hasLower || !hasNumber || !hasSpecial {
		return ErrInvalidPassword
	}

	return nil
}

// ValidateEmail checks if the email is valid
func (u *User) ValidateEmail() error {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(u.Email) {
		return ErrInvalidEmail
	}
	return nil
}

// ValidateUsername checks if the username is valid
func (u *User) ValidateUsername() error {
	usernameRegex := regexp.MustCompile(`^[a-zA-Z0-9_]{3,50}$`)
	if !usernameRegex.MatchString(u.Username) {
		return ErrInvalidUsername
	}
	return nil
}

func (u *User) ToResponse() dto.UserResponse {
	return dto.UserResponse{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
	}
}

func UsersToResponse(users []User) []dto.UserResponse {
	responses := make([]dto.UserResponse, len(users))
	for i, user := range users {
		responses[i] = user.ToResponse()
	}
	return responses
}

// Custom errors
var (
	ErrInvalidPassword     = &ValidationError{Message: "Password must be at least 8 characters long and contain uppercase, lowercase, number, and special character"}
	ErrInvalidEmail        = &ValidationError{Message: "Invalid email format"}
	ErrInvalidUsername     = &ValidationError{Message: "Username must be 3-50 characters long and contain only letters, numbers, and underscores"}
	ErrEmailExists         = &ValidationError{Message: "Email already registered"}
	ErrInvalidCredentials  = &ValidationError{Message: "Invalid email or password"}
	ErrForbidden          = &ValidationError{Message: "You don't have permission to perform this action"}
)

type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}