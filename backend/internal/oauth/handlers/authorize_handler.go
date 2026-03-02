package handlers

import (
	"net/http"
	"time"

	"terra/internal/middleware"
	"terra/internal/oauth/authcode"
	"terra/internal/oauth/authorization"
	"terra/internal/oauth/client"

	"github.com/google/uuid"
)

type AuthorizeHandler struct {
	authorizationRepo *authorization.Repository
	authCodeRepo      *authcode.Repository
	clientRepo        *client.Repository
}

func NewAuthorizeHandler(a *authorization.Repository, ac *authcode.Repository, c *client.Repository) *AuthorizeHandler {
	return &AuthorizeHandler{authorizationRepo: a, authCodeRepo: ac, clientRepo: c}
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
		http.Error(w, "Not logged in", 401)
		return
	}

	clientDBID, err := h.clientRepo.GetClientDBIDByClientID(r.Context(), clientID)

	if err != nil {
		http.Error(w, "server error", 500)
		return
	}

	exists, err := h.authorizationRepo.CheckExists(r.Context(), userID, clientDBID)

	if err != nil {
		http.Error(w, "server error", 500)
		return
	}

	if exists {
		code := uuid.NewString()

		authCode := &authcode.AuthorizationCode{
			Code:        code,
			ClientID:    clientID,
			UserID:      userID,
			RedirectURI: redirectURI,
			ExpiresAt:   time.Now().Add(5 * time.Minute),
		}

		err = h.authCodeRepo.Create(r.Context(), authCode)

		if err != nil {
			http.Error(w, "server error", 500)
			return
		}

		redirect := redirectURI + "?code=" + code + "&state=" + state

		http.Redirect(w, r, redirect, http.StatusFound)
		return
	}
}
