package handlers

import (
	"encoding/json"
	"net/http"
	"terra/internal/oauth/authcode"
	"terra/internal/oauth/client"
	"terra/internal/oauth/token"
	"time"

	"github.com/google/uuid"
)

type TokenHandler struct {
	authRepo   *authcode.Repository
	tokenRepo  *token.Repository
	clientRepo *client.Repository
}

func NewTokenHandler(a *authcode.Repository, t *token.Repository, c *client.Repository) *TokenHandler {
	return &TokenHandler{
		authRepo:   a,
		tokenRepo:  t,
		clientRepo: c,
	}
}

func (h *TokenHandler) Exchange(w http.ResponseWriter, r *http.Request) {
	clientID := r.FormValue("client_id")
	clientSecret := r.FormValue("client_secret")

	if clientID == "" || clientSecret == "" {
		http.Error(w, "client credential(s) missing", 400)
		return
	}

	client, err := h.clientRepo.FindByClientID(r.Context(), clientID)

	if err != nil {
		http.Error(w, "invalid client", 401)
		return
	}

	if client.ClientSecret != clientSecret {
		http.Error(w, "invalid client secret", 401)
		return
	}

	code := r.FormValue("code")

	if code == "" {
		http.Error(w, "code missing", 400)
		return
	}

	authCode, err := h.authRepo.Get(r.Context(), code)

	if authCode.ClientID != clientID {
		http.Error(w, "code not valid for this client", 401)
		return
	}

	if err != nil {
		http.Error(w, "invalid code", 400)
		return
	}

	if time.Now().After(authCode.ExpiresAt) {
		http.Error(w, "code expired", 400)
		return
	}

	err = h.authRepo.Delete(r.Context(), code)
	if err != nil {
		http.Error(w, "server error", 500)
		return
	}

	accessToken := &token.AccessToken{
		Token:     uuid.NewString(),
		UserID:    authCode.UserID,
		ClientID:  authCode.ClientID,
		ExpiresAt: time.Now().Add(1 * time.Hour),
	}

	err = h.tokenRepo.Create(r.Context(), accessToken)
	if err != nil {
		http.Error(w, "server error", 500)
		return
	}

	response := map[string]interface{}{
		"access_token": accessToken.Token,
		"token_type":   "Bearer",
		"expires_in":   3600,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
