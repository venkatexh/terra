package middleware

import (
	"context"
	"net/http"
	"time"

	"terra/internal/session"
)

type SessionMiddleware struct {
	sessionRepo *session.Repository
}

func NewSessionMiddleware(repo *session.Repository) *SessionMiddleware {
	return &SessionMiddleware{repo}
}


func (m *SessionMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("terra_session")
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		sess, err := m.sessionRepo.FindByID(r.Context(), cookie.Value)
		if err != nil || sess == nil {
			next.ServeHTTP(w, r)
			return
		}

		if time.Now().After(sess.ExpiresAt) {
			next.ServeHTTP(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), UserKey, sess.UserID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}