package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type dbEnv struct {
	User     string
	Password string
}

var singleton *dbEnv

func DBenv() *dbEnv {
	if singleton == nil {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		return &dbEnv{
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
		}
	}

	return singleton
}
