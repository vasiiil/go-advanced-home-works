package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DbConfig struct {
	Dsn string
}

type AuthConfig struct{
	Secret string
}

type EmailConfig struct {
	Email string
	Password string
	Address string
}

type Config struct {
	Db   DbConfig
	Auth AuthConfig
	Email EmailConfig
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using default config")
	}
	return &Config{
		Db: DbConfig{
			Dsn: os.Getenv("DSN"),
		},
		Auth: AuthConfig{
			Secret: os.Getenv("TOKEN"),
		},
		Email: EmailConfig{
			Email: os.Getenv("EMAIL"),
			Password: os.Getenv("EMAIL_PASSWORD"),
			Address: os.Getenv("EMAIL_ADDRESS"),
		},
	}
}
