package utils

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/joho/godotenv"
)

func TestLoadEnv(t *testing.T) {
	err := godotenv.Load(filepath.Join("..", ".env"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func TestGetEnv(t *testing.T) {
	log.Println(os.Getenv("JWT_ACCESS_SECRET"))
}
