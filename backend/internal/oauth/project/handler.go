package project

import (
	"encoding/json"
	"net/http"

	"terra/internal/middleware"

	"github.com/google/uuid"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{s}
}

func (h *Handler) CreateProject(w http.ResponseWriter, r *http.Request) {
	var req CreateProjectRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	name := req.Name
	description := req.Description

	project := &Project{
		ID:          uuid.New().String(),
		UserID:      r.Context().Value(middleware.UserKey).(string),
		Name:        name,
		Description: description,
	}

	if err := h.service.CreateProject(r.Context(), project); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetProjects(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserKey).(string)

	projects, err := h.service.FindByUserID(r.Context(), userID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}
