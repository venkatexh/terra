package db

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPool(databaseUrl string) *pgxpool.Pool {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	config, err := pgxpool.ParseConfig(databaseUrl)
	if err != nil {
		log.Fatal("Unable to parse DB config:", err)
	}

	config.MaxConns = 10
	config.MinConns = 2
	config.MaxConnLifetime = time.Hour

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatal("Unable to create DB pool:", err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		log.Fatal("Unable to ping DB:", err)
	}

	log.Println(("Connected to Postgres.."))

	return pool
}