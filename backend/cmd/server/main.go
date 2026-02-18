package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"terra/internal/auth"
	"terra/internal/config"
	"terra/internal/db"
	"terra/internal/oauth/authcode"
	"terra/internal/oauth/client"
	"terra/internal/oauth/handlers"
	oauthToken "terra/internal/oauth/token"
	"terra/internal/auth/token"
	"terra/internal/auth/user"
	"terra/internal/session"
	"terra/internal/middleware"
	authHandlers "terra/internal/auth/handlers"
	"terra/internal/auth/magic"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	// "github.com/google/uuid"
)

func main() {

	cfg := config.Load()

	pool := db.NewPool(cfg.DBUrl)
	defer pool.Close()

	userRepo := user.NewRepository(pool)

	// testUser := &user.User{
	// 	ID:    uuid.NewString(),
	// 	Email: "test@terra.dev",
	// }

	// err := userRepo.Create(context.Background(), testUser)

	// if err != nil {
	// 	log.Println("User insert error:", err)
	// } else {
	// 	log.Println("User created:", testUser)
	// }

	authRepo := authcode.NewRepository(pool)
	oauthTokenRepo := oauthToken.NewRepository(pool)
	tokenRepo := token.NewRepository(pool)
	clientRepo := client.NewRepository(pool)
	sessionRepo := session.NewRepository(pool)

	authorizeHandler := handlers.NewAuthorizeHandler(authRepo)
	tokenHandler := handlers.NewTokenHandler(authRepo, oauthTokenRepo, clientRepo)
	loginHandler := authHandlers.NewLoginHandler(sessionRepo)

	magicService := magic.NewService(userRepo, tokenRepo, sessionRepo, magic.ConsoleMailer{})
	magicHandler := magic.NewHandler(magicService)

	// raw, hash, _ := token.GenerateToken()

	// loginToken := &token.LoginToken{
	// 	ID:        uuid.NewString(),
	// 	UserID:    testUser.ID,
	// 	Hash:      hash,
	// 	ExpiresAt: time.Now().Add(time.Hour),
	// }

	// err = tokenRepo.Create(context.Background(), loginToken)
	// if err != nil {
	// 	log.Println("Login token insert error:", err)
	// } else {
	// 	log.Println("Login token created:", loginToken)
	// 	log.Println("Raw token (for email):", raw)
	// }

	authSvc := auth.NewService(userRepo, tokenRepo)

	clientSvc := client.NewService(clientRepo)

	client, err := clientSvc.RegisterClient(
		context.Background(),
		"test",
		[]string{"http://localhost:3000/callback"},
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("ClientID:", client.ClientID)
	log.Println("ClientSecret:", client.ClientSecret)

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "PUT", "POST", "DELETE"},
		AllowedHeaders:   []string{"Access", "Authorization", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Use(middleware.NewSessionMiddleware(sessionRepo).Middleware)

	r.Get("/login", loginHandler.Login)

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

	r.Post("/auth/request-link", auth.RequestLinkHandler(authSvc))

	r.Post("/auth/verify-link", auth.VerifyLinkHandler(authSvc))

	r.Get("/oauth/authorize", authorizeHandler.Authorize)

	r.Post("/oauth/token", tokenHandler.Exchange)

	r.Post("/auth/magic-link/request", magicHandler.RequestLink)
	r.Get("/auth/magic-link/verify", magicHandler.VerifyLink)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Terra auth server starting on port 8080..")
	log.Fatal(server.ListenAndServe())
}
