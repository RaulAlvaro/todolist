package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBDSN      string // Cadena de conexión
	ServerPort string
}

func Load() *Config {
	_ = godotenv.Load()
	// En lugar de initializers, cargamos aquí
	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		// DSN por defecto para desarrollo
		dsn = "host=localhost user=postgres password=postgres dbname=todolist port=5432 sslmode=disable"
	}

	return &Config{
		DBDSN:      dsn,
		ServerPort: os.Getenv("PORT"),
	}
}
