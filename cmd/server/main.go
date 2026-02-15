package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"terra/internal/config"
	"terra/internal/db"
	"terra/internal/user"
	"terra/internal/token"
	"terra/internal/auth"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/google/uuid"
)

func main() {

	cfg := config.Load()

	pool := db.NewPool(cfg.DBUrl)
	defer pool.Close()

	userRepo := user.NewRepository(pool)

	testUser := &user.User{
		ID: uuid.NewString(),
		Email: "test@terra.dev",
	}

	err := userRepo.Create(context.Background(), testUser)

	if err != nil {
		log.Println("User insert error:", err)
	} else {
		log.Println("User created:", testUser)
	}

	tokenRepo := token.NewRepository(pool)

	raw, hash, _ := token.GenerateToken()
	
	loginToken := &token.LoginToken{
		ID: uuid.NewString(),
		UserID: testUser.ID,
		Hash: hash,
		ExpiresAt: time.Now().Add(time.Hour),
	}

	err = tokenRepo.Create(context.Background(), loginToken)
	if err != nil {
		log.Println("Login token insert error:", err)
	} else {
		log.Println("Login token created:", loginToken)
		log.Println("Raw token (for email):", raw)
	}

	authSvc := auth.NewService(userRepo, tokenRepo)


	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET","PUT","POST","DELETE"},
		AllowedHeaders: []string{"Access","Authorization","Content-Type"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Terra server running.."))
	})

	r.Post("/auth/request-link", auth.RequestLinkHandler(authSvc))

	r.Post("/auth/verify-link", auth.VerifyLinkHandler(authSvc))

	server := &http.Server{
		Addr: ":8080",
		Handler: r,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Terra auth server starting on port 8080..")
	log.Fatal(server.ListenAndServe())
}