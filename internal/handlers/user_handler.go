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

type UserHandler struct {
	repo *repository.UserRepository
}

func NewUserHandler(userRepository *repository.UserRepository) *UserHandler {
	return &UserHandler{
		repo: userRepository,
	}
}

func (h *UserHandler) GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		dto.WriteError(w, http.StatusBadRequest, err)
		return
	}

	user, err := h.repo.GetByID(uint(userID))
	if err != nil {
		dto.WriteError(w, http.StatusNotFound, err)
		return
	}

	dto.WriteSuccess(w, http.StatusOK, dto.UserToResponse(user), "User retrieved successfully")
}

func (h *UserHandler) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
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

	// Get paginated users
	users, total, err := h.repo.GetAllUsersPaginated(page, pageSize)
	if err != nil {
		dto.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	response := dto.PaginatedResponse{
		Response: dto.NewSuccessResponse(dto.UsersToResponse(users), "Users retrieved successfully"),
		Page:     page,
		PageSize: pageSize,
		Total:    int(total),
	}

	dto.WriteJSON(w, http.StatusOK, response)
}

func (h *UserHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if !middleware.IsUserSelf(r, vars["id"]) {
		dto.WriteError(w, http.StatusForbidden, models.ErrForbidden)
		return
	}

	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		dto.WriteError(w, http.StatusBadRequest, err)
		return
	}

	var updateReq dto.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&updateReq); err != nil {
		dto.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Get existing user
	user, err := h.repo.GetByID(uint(userID))
	if err != nil {
		dto.WriteError(w, http.StatusNotFound, err)
		return
	}

	// Update fields if provided
	if updateReq.Username != "" {
		user.Username = updateReq.Username
	}
	if updateReq.Email != "" {
		user.Email = updateReq.Email
	}
	if updateReq.Password != "" {
		user.Password = updateReq.Password
	}

	// Validate updated fields
	if err := user.ValidateUsername(); err != nil {
		dto.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if err := user.ValidateEmail(); err != nil {
		dto.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if updateReq.Password != "" {
		if err := user.ValidatePassword(); err != nil {
			dto.WriteError(w, http.StatusBadRequest, err)
			return
		}
	}

	if err := h.repo.UpdateUser(user); err != nil {
		dto.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	dto.WriteSuccess(w, http.StatusOK, dto.UserToResponse(user), "User updated successfully")
}

func (h *UserHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if !middleware.IsUserSelf(r, vars["id"]) {
		dto.WriteError(w, http.StatusForbidden, models.ErrForbidden)
		return
	}

	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		dto.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.repo.DeleteUser(uint(userID)); err != nil {
		dto.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	dto.WriteSuccess(w, http.StatusOK, nil, "User deleted successfully")
}
