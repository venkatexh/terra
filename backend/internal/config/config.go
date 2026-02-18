package config

import (
	"log"
	"os"
)

type Config struct {
	DBUrl string 
}

func Load() *Config {
	dbUrl := os.Getenv("DATABASE_URL")

	if dbUrl == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	return &Config {
		DBUrl: dbUrl,
	}
}