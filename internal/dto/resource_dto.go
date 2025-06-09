package dto

import (
	"devlink/internal/models"
	"encoding/json"
)

type ResourceResponse struct {
	ID          uint                `json:"id"`
	Title       string              `json:"title"`
	Type        models.ResourceType `json:"type"`
	URL         string              `json:"url,omitempty"`
	Category    models.LinkCategory `json:"category,omitempty"`
	Description string              `json:"description"`
	Tags        []string            `json:"tags"`
	Language    string              `json:"language,omitempty"`
	CodeContent string              `json:"code_content,omitempty"`
	UserID      uint                `json:"user_id"`
}

type CreateResourceRequest struct {
	Title       string              `json:"title" validate:"required,min=3,max=100"`
	Type        models.ResourceType `json:"type" validate:"required,oneof=link code"`
	URL         string              `json:"url" validate:"omitempty,url"`
	Category    models.LinkCategory `json:"category" validate:"omitempty,oneof=github article tool other"`
	Description string              `json:"description" validate:"max=500"`
	Tags        []string            `json:"tags" validate:"max=10,dive,max=30"`
	Language    string              `json:"language" validate:"omitempty,min=2,max=20"`
	CodeContent string              `json:"code_content" validate:"omitempty,min=1,max=10000"`
}

type UpdateResourceRequest struct {
	Title       string              `json:"title" validate:"omitempty,min=3,max=100"`
	Type        models.ResourceType `json:"type" validate:"omitempty,oneof=link code"`
	URL         string              `json:"url" validate:"omitempty,url"`
	Category    models.LinkCategory `json:"category" validate:"omitempty,oneof=github article tool other"`
	Description string              `json:"description" validate:"omitempty,max=500"`
	Tags        []string            `json:"tags" validate:"omitempty,max=10,dive,max=30"`
	Language    string              `json:"language" validate:"omitempty,min=2,max=20"`
	CodeContent string              `json:"code_content" validate:"omitempty,min=1,max=10000"`
}

func ResourceToResponse(resource *models.Resource) ResourceResponse {
	var tags []string
	if resource.Tags != nil {
		json.Unmarshal(resource.Tags, &tags)
	}

	return ResourceResponse{
		ID:          resource.ID,
		Title:       resource.Title,
		Type:        resource.Type,
		URL:         resource.URL,
		Category:    resource.Category,
		Description: resource.Description,
		Tags:        tags,
		Language:    resource.Language,
		CodeContent: resource.CodeContent,
		UserID:      resource.UserID,
	}
}

func ResourcesToResponse(resources []models.Resource) []ResourceResponse {
	responses := make([]ResourceResponse, len(resources))
	for i, resource := range resources {
		responses[i] = ResourceToResponse(&resource)
	}
	return responses
}
