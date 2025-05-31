package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	// Path to the .env file
	EnvPath = ".env"

	// Error message when .env file is not found
	NoEnvPathError = "no .env file found, please create one from .env.example"
)

// Load environment variables
func init() {
	err := godotenv.Load(EnvPath)
	if err != nil {
		log.Fatal(NoEnvPathError)
	}
}

// Get environment value by key
func Get(key EnvKey) string {
	// Validate key
	if !key.IsValid() {
		return ""
	}

	return os.Getenv(string(key))
}
