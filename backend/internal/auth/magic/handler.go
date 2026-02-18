package magic

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{s}
}

type requestBody struct {
	Email string `json:"email"`
}

func (h *Handler) RequestLink(w http.ResponseWriter, r *http.Request) {
	var body requestBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if err := h.service.RequestMagicLink(r.Context(), body.Email); err != nil {
		log.Println("MAGIC LINK ERROR:", err)
		http.Error(w, "failed to send link", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) VerifyLink(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")

	if token == "" {
		http.Error(w, "missing token", http.StatusBadRequest)
		return
	}

	sessionID, err := h.service.VerifyMagicToken(r.Context(), token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "terra_session",
		Value:    sessionID,
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // true in production
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().Add(7 * 24 * time.Hour),
	})

	w.Write([]byte("Login successful"))
}

