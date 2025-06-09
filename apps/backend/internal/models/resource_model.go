package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type ResourceType string
type LinkCategory string

const (
	ResourceTypeLink ResourceType = "link"
	ResourceTypeCode ResourceType = "code"
)

const (
	LinkCategoryGitHub  LinkCategory = "github"
	LinkCategoryArticle LinkCategory = "article"
	LinkCategoryTool    LinkCategory = "tool"
	LinkCategoryOther   LinkCategory = "other"
)

type Resource struct {
	gorm.Model
	Title       string         `json:"title" gorm:"not null"`
	Type        ResourceType   `json:"type" gorm:"not null;type:varchar(10)"`
	URL         string         `json:"url" gorm:"uniqueIndex"`
	Category    LinkCategory   `json:"category" gorm:"type:varchar(20)"`
	Description string         `json:"description"`
	Tags        datatypes.JSON `json:"tags"`

	// Code snippet specific fields
	Language    string `json:"language"`
	CodeContent string `json:"code_content" gorm:"type:text"`

	UserID uint `json:"user_id" gorm:"not null"`
}

// Validate checks if the resource is valid based on its type
func (r *Resource) Validate() error {
	switch r.Type {
	case ResourceTypeLink:
		if r.URL == "" {
			return &ValidationError{Message: "URL is required for link resources"}
		}
		if r.Category == "" {
			return &ValidationError{Message: "Category is required for link resources"}
		}
	case ResourceTypeCode:
		if r.CodeContent == "" {
			return &ValidationError{Message: "Code content is required for code resources"}
		}
		if r.Language == "" {
			return &ValidationError{Message: "Language is required for code resources"}
		}
	default:
		return &ValidationError{Message: "Invalid resource type"}
	}
	return nil
}
