package client

import (
	"encoding/json"
	"net/http"
	"terra/internal/middleware"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{s}
}

func (h *Handler) CreateClient(w http.ResponseWriter, r *http.Request) {

	projectID := chi.URLParam(r, "projectId")

	var req Client

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID := r.Context().Value(middleware.UserKey).(string)

	client, err := h.service.RegisterClient(r.Context(), req.Name, req.RedirectURIs, projectID, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(client)
}

func (h *Handler) GetClients(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value(middleware.UserKey).(string)
	projectID := chi.URLParam(r, "projectId")

	clients, err := h.service.FindByUserID(r.Context(), userID, projectID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clients)
}

