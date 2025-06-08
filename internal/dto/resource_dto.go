package dto

type ResourceResponse struct {
	ID          uint     `json:"id"`
	Title       string   `json:"title"`
	URL         string   `json:"url"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	UserID      uint     `json:"user_id"`
}

type CreateResourceRequest struct {
	Title       string   `json:"title" validate:"required,min=3,max=100"`
	URL         string   `json:"url" validate:"required,url"`
	Description string   `json:"description" validate:"max=500"`
	Tags        []string `json:"tags" validate:"max=10,dive,max=30"`
}

type UpdateResourceRequest struct {
	Title       string   `json:"title" validate:"omitempty,min=3,max=100"`
	URL         string   `json:"url" validate:"omitempty,url"`
	Description string   `json:"description" validate:"omitempty,max=500"`
	Tags        []string `json:"tags" validate:"omitempty,max=10,dive,max=30"`
}

func ResourceToResponse(resource *models.Resource) ResourceResponse {
	return ResourceResponse{
		ID:          resource.ID,
		Title:       resource.Title,
		URL:         resource.URL,
		Description: resource.Description,
		Tags:        resource.Tags,
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