package config

type EnvKey string

// Env params keys
const (
	EnvAppAddress EnvKey = "APP_ADDRESS" // e.g. "localhost:8080"
	EnvDBHost     EnvKey = "DB_HOST"     // e.g. "localhost"
	EnvDBPort     EnvKey = "DB_PORT"     // e.g. "5432"
	EnvDBUser     EnvKey = "DB_USER"     // e.g. "postgres"
	EnvDBPassword EnvKey = "DB_PASSWORD" // e.g. "password"
	EnvDBName     EnvKey = "DB_NAME"     // e.g. "tasks_db"
)

// Check if the environment key is valid
func (e EnvKey) IsValid() bool {
	switch e {
	case EnvAppAddress, EnvDBHost, EnvDBPort, EnvDBUser,
		EnvDBPassword, EnvDBName:
		return true
	default:
		return false
	}
}
