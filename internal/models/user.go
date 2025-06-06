package models

import (
	"devlink/internal/dto"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string     `json:"username" gorm:"not null;uniqueIndex"`
	Email     string     `json:"email" gorm:"not null;uniqueIndex"`
	Password  string     `json:"password" gorm:"not null"`
	Resources []Resource `json:"resources"`
}

func (u *User) ToResponse() dto.UserResponse {
	return dto.UserResponse{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
	}
}
