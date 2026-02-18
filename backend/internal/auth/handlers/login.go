package handlers

import (
	"net/http"
	"time"

	"github.com/google/uuid"

	"terra/internal/session"
)

type LoginHandler struct {
	sessionRepo *session.Repository
}

func NewLoginHandler(repo *session.Repository) *LoginHandler {
	return &LoginHandler{sessionRepo: repo}
}

func (h *LoginHandler) Login(w http.ResponseWriter, r *http.Request) {
	userID := "user-1"

	sessionID := uuid.NewString()

	s  := &session.Session{
		ID: sessionID,
		UserID: userID,
		ExpiresAt: time.Now().Add(30 * 24 * time.Hour),
	}

	err := h.sessionRepo.Create(r.Context(), s)

	if err != nil {
		http.Error(w, "server error", 500)
		return
	}

	http.SetCookie(w, &http.Cookie {
		Name: "terra_session",
		Value: sessionID,
		Path: "/",
		HttpOnly: true,
	})

	w.Write([]byte("Logged in!"))
}