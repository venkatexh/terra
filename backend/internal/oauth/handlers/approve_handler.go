package handlers

import (
	"net/http"
	"time"

	"terra/internal/middleware"

	"terra/internal/oauth/authcode"
	"terra/internal/oauth/authorization"
	"terra/internal/oauth/client"
	"terra/internal/oauth/project"
	"terra/internal/oauth/token"

	"github.com/google/uuid"
)

type ApproveHandler struct {
	authCodeRepo      *authcode.Repository
	tokenRepo         *token.Repository
	clientRepo        *client.Repository
	projectRepo       *project.Repository
	authorizationRepo *authorization.Repository
}

func NewApproveHandler(ac *authcode.Repository, t *token.Repository, c *client.Repository, p *project.Repository, a *authorization.Repository) *ApproveHandler {
	return &ApproveHandler{
		authCodeRepo:      ac,
		tokenRepo:         t,
		clientRepo:        c,
		projectRepo:       p,
		authorizationRepo: a,
	}
}

func (h *ApproveHandler) Approve(w http.ResponseWriter, r *http.Request) {

	state := r.URL.Query().Get("state")
	groupID := r.URL.Query().Get("group_id")
	clientID := r.URL.Query().Get("client_id")
	redirectURI := r.URL.Query().Get("redirect_uri")
	userID, ok := r.Context().Value(middleware.UserKey).(string)

	if !ok {
		http.Error(w, "Not logged in", http.StatusUnauthorized)
		return
	}

	if clientID == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
	}

	clientDBID, err := h.clientRepo.GetClientDBIDByClientID(r.Context(), clientID)

	if err != nil {
		http.Error(w, "server error", 500)
		return
	}

	err = h.authorizationRepo.Create(r.Context(), groupID, clientDBID)

	if err != nil {
		http.Error(w, "server error", 500)
		return
	}

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
}
