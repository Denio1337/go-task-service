package config

import "errors"

type EnvKey string

// Path to .env file
const EnvPath = ".env"

// Environment variable keys
const (
	EnvAppAddress EnvKey = "APP_ADDRESS" // e.g. "localhost:8080"

	EnvDBHost     EnvKey = "DB_HOST"     // e.g. "localhost"
	EnvDBPort     EnvKey = "DB_PORT"     // e.g. "5432"
	EnvDBUser     EnvKey = "DB_USER"     // e.g. "postgres"
	EnvDBPassword EnvKey = "DB_PASSWORD" // e.g. "password"
	EnvDBName     EnvKey = "DB_NAME"     // e.g. "tasks_db"

	EnvMultipleProcesses EnvKey = "MULTIPLE_PROCESSES" // e.g. "true" or "false"
)

// Env file was not found
var ErrEnvNotFound = errors.New(".env file not found")

// Check if the environment key is valid
func (e EnvKey) IsValid() bool {
	switch e {
	case EnvAppAddress, EnvDBHost, EnvDBPort, EnvDBUser,
		EnvDBPassword, EnvDBName, EnvMultipleProcesses:
		return true
	default:
		return false
	}
}
