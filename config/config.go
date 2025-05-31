package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Load environment variables
func init() {
	err := godotenv.Load(EnvPath)
	if err != nil {
		log.Fatal(ErrEnvNotFound)
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

func GetBool(key EnvKey) bool {
	value := Get(key)
	if value == "" {
		return false
	}

	switch value {
	case "true", "1":
		return true
	default:
		return false
	}
}
