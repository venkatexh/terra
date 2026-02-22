package authorization

import (
	"encoding/json"
	"log"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{s}
}

func (h *Handler) CreateAuthorization(w http.ResponseWriter, r *http.Request) {
	var req AuthorizationRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	groupID := req.GroupID
	clientDBID := req.ClientDBID

	if err := h.service.CreateAuthorization(r.Context(), groupID, clientDBID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetAuthorizationByGroupID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("groupId")

	authorization, err := h.service.FindByGroupID(r.Context(), id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(authorization)
}