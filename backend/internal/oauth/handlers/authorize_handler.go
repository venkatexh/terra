package handlers

import (
	"net/http"
	"time"

	"github.com/google/uuid"

	"terra/internal/middleware"
	"terra/internal/oauth/authcode"
)

type AuthorizeHandler struct {
	authRepo *authcode.Repository
}

func NewAuthorizeHandler(repo *authcode.Repository) *AuthorizeHandler {
	return &AuthorizeHandler{authRepo: repo}
}

func (h *AuthorizeHandler) Authorize(w http.ResponseWriter, r *http.Request) {

	clientID := r.URL.Query().Get("client_id")
	redirectURI := r.URL.Query().Get("redirect_uri")
	state := r.URL.Query().Get("state")

	if clientID == "" || redirectURI == "" {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	userID, ok := r.Context().Value(middleware.UserKey).(string)

	if !ok {
		http.Error(w, "not logged in", http.StatusUnauthorized)
		return
	}

	code := uuid.NewString()

	authCode := &authcode.AuthorizationCode{
		Code: code,
		ClientID: clientID,
		UserID: userID,
		RedirectURI: redirectURI,
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}

	err := h.authRepo.Create(r.Context(), authCode)

	if err != nil {
		http.Error(w, "server error", 500)
		return
	}

	redirect := redirectURI + "?code=" + code + "&state=" + state

	http.Redirect(w, r, redirect, http.StatusFound)
}