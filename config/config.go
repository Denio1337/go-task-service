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

// Load .env file
func init() {
	err := godotenv.Load(EnvPath)
	if err != nil {
		log.Fatal(NoEnvPathError)
	}
}

// Get .env value by key
func Get(key EnvKey) string {
	// Return default value "" if env key is invalid
	if !key.IsValid() {
		return ""
	}

	return os.Getenv(string(key))
}
