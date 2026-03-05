package otp

import (
	"encoding/json"
	"log"
	"net"
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

func (h *Handler) RequestOTP(w http.ResponseWriter, r *http.Request) {
	var body requestBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if err := h.service.RequestOTP(r.Context(), body.Email); err != nil {
		log.Println("MAGIC LINK ERROR:", err)
		http.Error(w, "failed to send link", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) VerifyOTP(w http.ResponseWriter, r *http.Request) {

	token := r.URL.Query().Get("otp")

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return
	}

	ua := r.UserAgent()

	if token == "" {
		http.Error(w, "Missing OTP", http.StatusBadRequest)
		return
	}

	sessionID, err := h.service.VerifyOTP(r.Context(), token, ip, ua)
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
		Expires:  time.Now().Add(30 * 24 * time.Hour),
	})

	w.Write([]byte("Login successful"))
}
