package handlers

import (
	"devlink/internal/dto"
	"devlink/internal/middleware"
	"devlink/internal/models"
	"devlink/internal/repository"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ResourceHandler struct {
	repo *repository.ResourceRepository
}

func NewResourceHandler(resourceRepository *repository.ResourceRepository) *ResourceHandler {
	return &ResourceHandler{
		repo: resourceRepository,
	}
}

func (h *ResourceHandler) CreateResourceHandler(w http.ResponseWriter, r *http.Request) {
	var createReq dto.CreateResourceRequest
	if err := json.NewDecoder(r.Body).Decode(&createReq); err != nil {
		dto.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Get user ID from JWT
	claims, ok := middleware.GetUserClaims(r)
	if !ok {
		dto.WriteError(w, http.StatusUnauthorized, models.ErrInvalidCredentials)
		return
	}
	userID := uint(claims["user_id"].(float64))

	// Create resource
	resource := &models.Resource{
		Title:       createReq.Title,
		URL:         createReq.URL,
		Description: createReq.Description,
		Tags:        createReq.Tags,
		UserID:      userID,
	}

	if err := h.repo.CreateResource(resource); err != nil {
		dto.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	dto.WriteSuccess(w, http.StatusCreated, dto.ResourceToResponse(resource), "Resource created successfully")
}

func (h *ResourceHandler) GetResourceByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	resourceID, err := strconv.Atoi(vars["id"])
	if err != nil {
		dto.WriteError(w, http.StatusBadRequest, err)
		return
	}

	resource, err := h.repo.GetByID(uint(resourceID))
	if err != nil {
		dto.WriteError(w, http.StatusNotFound, err)
		return
	}

	// Check if user owns the resource
	claims, ok := middleware.GetUserClaims(r)
	if !ok {
		dto.WriteError(w, http.StatusUnauthorized, models.ErrInvalidCredentials)
		return
	}
	userID := uint(claims["user_id"].(float64))
	if resource.UserID != userID {
		dto.WriteError(w, http.StatusForbidden, models.ErrForbidden)
		return
	}

	dto.WriteSuccess(w, http.StatusOK, dto.ResourceToResponse(resource), "Resource retrieved successfully")
}

func (h *ResourceHandler) GetUserResourcesHandler(w http.ResponseWriter, r *http.Request) {
	// Get user ID from JWT
	claims, ok := middleware.GetUserClaims(r)
	if !ok {
		dto.WriteError(w, http.StatusUnauthorized, models.ErrInvalidCredentials)
		return
	}
	userID := uint(claims["user_id"].(float64))

	// Parse pagination parameters
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))

	// Set default values if not provided
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	// Get resources
	resources, total, err := h.repo.GetByUserID(userID, page, pageSize)
	if err != nil {
		dto.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	response := dto.PaginatedResponse{
		Response: dto.NewSuccessResponse(dto.ResourcesToResponse(resources), "Resources retrieved successfully"),
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	}

	dto.WriteJSON(w, http.StatusOK, response)
}

func (h *ResourceHandler) UpdateResourceHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	resourceID, err := strconv.Atoi(vars["id"])
	if err != nil {
		dto.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Get existing resource
	resource, err := h.repo.GetByID(uint(resourceID))
	if err != nil {
		dto.WriteError(w, http.StatusNotFound, err)
		return
	}

	// Check if user owns the resource
	claims, ok := middleware.GetUserClaims(r)
	if !ok {
		dto.WriteError(w, http.StatusUnauthorized, models.ErrInvalidCredentials)
		return
	}
	userID := uint(claims["user_id"].(float64))
	if resource.UserID != userID {
		dto.WriteError(w, http.StatusForbidden, models.ErrForbidden)
		return
	}

	// Parse update request
	var updateReq dto.UpdateResourceRequest
	if err := json.NewDecoder(r.Body).Decode(&updateReq); err != nil {
		dto.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Update fields if provided
	if updateReq.Title != "" {
		resource.Title = updateReq.Title
	}
	if updateReq.URL != "" {
		resource.URL = updateReq.URL
	}
	if updateReq.Description != "" {
		resource.Description = updateReq.Description
	}
	if updateReq.Tags != nil {
		resource.Tags = updateReq.Tags
	}

	if err := h.repo.UpdateResource(resource); err != nil {
		dto.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	dto.WriteSuccess(w, http.StatusOK, dto.ResourceToResponse(resource), "Resource updated successfully")
}

func (h *ResourceHandler) DeleteResourceHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	resourceID, err := strconv.Atoi(vars["id"])
	if err != nil {
		dto.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Get existing resource
	resource, err := h.repo.GetByID(uint(resourceID))
	if err != nil {
		dto.WriteError(w, http.StatusNotFound, err)
		return
	}

	// Check if user owns the resource
	claims, ok := middleware.GetUserClaims(r)
	if !ok {
		dto.WriteError(w, http.StatusUnauthorized, models.ErrInvalidCredentials)
		return
	}
	userID := uint(claims["user_id"].(float64))
	if resource.UserID != userID {
		dto.WriteError(w, http.StatusForbidden, models.ErrForbidden)
		return
	}

	if err := h.repo.DeleteResource(uint(resourceID)); err != nil {
		dto.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	dto.WriteSuccess(w, http.StatusOK, nil, "Resource deleted successfully")
}

func (h *ResourceHandler) SearchResourcesHandler(w http.ResponseWriter, r *http.Request) {
	// Get user ID from JWT
	claims, ok := middleware.GetUserClaims(r)
	if !ok {
		dto.WriteError(w, http.StatusUnauthorized, models.ErrInvalidCredentials)
		return
	}
	userID := uint(claims["user_id"].(float64))

	// Get search query
	query := r.URL.Query().Get("q")
	if query == "" {
		dto.WriteError(w, http.StatusBadRequest, models.ErrInvalidRequest)
		return
	}

	// Parse pagination parameters
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))

	// Set default values if not provided
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	// Search resources
	resources, total, err := h.repo.SearchResources(query, userID, page, pageSize)
	if err != nil {
		dto.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	response := dto.PaginatedResponse{
		Response: dto.NewSuccessResponse(dto.ResourcesToResponse(resources), "Resources retrieved successfully"),
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	}

	dto.WriteJSON(w, http.StatusOK, response)
}

func (h *ResourceHandler) GetResourcesByTagsHandler(w http.ResponseWriter, r *http.Request) {
	// Get user ID from JWT
	claims, ok := middleware.GetUserClaims(r)
	if !ok {
		dto.WriteError(w, http.StatusUnauthorized, models.ErrInvalidCredentials)
		return
	}
	userID := uint(claims["user_id"].(float64))

	// Get tags from query parameter
	tagsStr := r.URL.Query().Get("tags")
	if tagsStr == "" {
		dto.WriteError(w, http.StatusBadRequest, models.ErrInvalidRequest)
		return
	}

	// Parse tags
	var tags []string
	if err := json.Unmarshal([]byte(tagsStr), &tags); err != nil {
		dto.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Parse pagination parameters
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))

	// Set default values if not provided
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	// Get resources by tags
	resources, total, err := h.repo.GetByTags(tags, userID, page, pageSize)
	if err != nil {
		dto.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	response := dto.PaginatedResponse{
		Response: dto.NewSuccessResponse(dto.ResourcesToResponse(resources), "Resources retrieved successfully"),
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	}

	dto.WriteJSON(w, http.StatusOK, response)
} 