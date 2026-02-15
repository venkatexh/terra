package auth

import (
	"encoding/json"
	"net/http"
)

type requestLinkPayload struct {
	Email string `json:"email"`
}

func RequestLinkHandler(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload requestLinkPayload

		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		err = svc.RequestLoginLink(r.Context(), payload.Email)
		if err != nil {
			http.Error(w, "Failed to request login link", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Login link sent"))
	}
}

type verifyPayload struct {
	Token string `json:"token"`
}

func VerifyLinkHandler(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload verifyPayload

		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil || payload.Token == "" {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		jwtToken, err := svc.VerifyLoginToken(r.Context(), payload.Token)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"token": jwtToken,
		})
	}
}
