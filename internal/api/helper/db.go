package helper

import (
	"database/sql"
	"log"
	"os"

	"github.com/naijasonezeiru/go-phish-backend/internal/database"
)

type ApiConfig struct {
	DB *database.Queries
}

func ConnectDB() ApiConfig {
	dbURL := os.Getenv("POSTGRES_URL")
	if dbURL == "" {
		log.Fatal("POSTGRES_URL is not found in the environment")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't dbect to the database", err)
	}

	apiCfg := ApiConfig{
		DB: database.New(db),
	}

	return apiCfg

}
