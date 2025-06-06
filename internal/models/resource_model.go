package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Resource struct {
	gorm.Model
	Title       string         `json:"title" gorm:"not null"`
	URL         string         `json:"url" gorm:"not null;uniqueIndex"`
	Description string         `json:"description"`
	Tags        datatypes.JSON `json:"tags"`

	UserID uint `json:"user_id" gorm:"not null"`
}
