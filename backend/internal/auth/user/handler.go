package user

import (
	"encoding/json"
	"net/http"
	"terra/internal/middleware"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{s}
}

func (h *Handler) FindMe(w http.ResponseWriter, r *http.Request) {

	loggedInUserID := r.Context().Value(middleware.UserKey).(string)

	user, err := h.service.Me(r.Context(), loggedInUserID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
