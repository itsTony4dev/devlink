package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"not null;uniqueIndex"`
	Email    string `json:"email" gorm:"not null;uniqueIndex"`
	Password string `json:"password" gorm:"not null"`
	Resources []Resource `json:"resources"`
}