package main

import (
	// "context"
	"log"
	"net/http"
	"terra/internal/auth/magic"
	"terra/internal/auth/token"
	"terra/internal/auth/user"
	"terra/internal/config"
	"terra/internal/db"
	"terra/internal/email"
	"terra/internal/group"
	"terra/internal/middleware"
	"terra/internal/oauth/authcode"
	"terra/internal/oauth/authorization"
	"terra/internal/oauth/client"
	"terra/internal/oauth/handlers"
	"terra/internal/oauth/project"
	oauthToken "terra/internal/oauth/token"
	"terra/internal/session"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {

	cfg := config.Load()

	pool := db.NewPool(cfg.DBUrl)
	defer pool.Close()

	userRepo := user.NewRepository(pool)
	clientRepo := client.NewRepository(pool)
	authRepo := authcode.NewRepository(pool)
	oauthTokenRepo := oauthToken.NewRepository(pool)
	tokenRepo := token.NewRepository(pool)
	sessionRepo := session.NewRepository(pool)
	groupRepo := group.NewRepository(pool)
	authorizationRepo := authorization.NewRepository(pool)
	projectRepo := project.NewRepository(pool)

	authorizeHandler := handlers.NewAuthorizeHandler(authRepo)
	tokenHandler := handlers.NewTokenHandler(authRepo, oauthTokenRepo, clientRepo)
	// loginHandler := authHandlers.NewLoginHandler(sessionRepo)

	clientSvc := client.NewService(clientRepo, projectRepo)
	// authSvc := auth.NewService(userRepo, tokenRepo)

	emailSvc := email.NewService()
	magicService := magic.NewService(userRepo, tokenRepo, sessionRepo, emailSvc)
	groupSvc := group.NewService(groupRepo)
	authorizationSvc := authorization.NewService(authorizationRepo)
	projectsSvc := project.NewService(projectRepo)

	magicHandler := magic.NewHandler(magicService)
	groupHandler := group.NewHandler(groupSvc)
	authorizationHandler := authorization.NewHandler(authorizationSvc)
	projectHandler := project.NewHandler(projectsSvc)
	clientHandler := client.NewHandler(clientSvc)

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "PUT", "POST", "DELETE"},
		AllowedHeaders: []string{
			"Access",
			"Authorization",
			"Content-Type",
			"X-CSRF-Token",
		},

		ExposedHeaders: []string{
			"Link",
		},

		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Use(middleware.NewSessionMiddleware(sessionRepo).Middleware)

	// r.Get("/login", loginHandler.Login)

	r.Get("/me", func(w http.ResponseWriter, r *http.Request) {
		userID, ok := r.Context().Value(middleware.UserKey).(string)

		if !ok {
			w.Write([]byte("Not logged in"))
			return
		}

		w.Write([]byte("Logged in user with ID: " + userID))
	})

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Terra server running.."))
	})

	// r.Post("/auth/request-link", auth.RequestLinkHandler(authSvc))

	// r.Post("/auth/verify-link", auth.VerifyLinkHandler(authSvc))

	r.Get("/oauth/authorize", authorizeHandler.Authorize)

	r.Post("/oauth/token", tokenHandler.Exchange)

	r.Post("/auth/magic-link/request", magicHandler.RequestLink)
	r.Get("/auth/magic-link/verify", magicHandler.VerifyLink)

	r.Post("/authorizations", authorizationHandler.CreateAuthorization)
	r.Get("/authorizations", authorizationHandler.GetAuthorizationByGroupID)

	r.Get("/me/groups", groupHandler.GetGroups)
	r.Get("/me/groups/{appGroupID}", groupHandler.GetGroup)

	r.Post("/projects", projectHandler.CreateProject)
	r.Get("/me/projects", projectHandler.GetProjects)
	r.Get("/projects/{projectId}", projectHandler.GetProject)

	r.Post("/projects/{projectId}/clients", clientHandler.CreateClient)
	r.Get("/projects/{projectId}/clients", clientHandler.GetClients)

	r.Get("/oauth/clients/{clientId}", clientHandler.GetClient)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Terra auth server starting on port 8080..")
	log.Fatal(server.ListenAndServe())
}
