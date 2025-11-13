package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DbConfig struct {
	Dsn string
}

type AuthConfig struct {
	JwtSecret string
}

type SmsConfig struct {
	AccountSid string
	AuthToken  string
}

type Config struct {
	Db   DbConfig
	Auth AuthConfig
	Sms  SmsConfig
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
			JwtSecret: os.Getenv("JWT_SECRET"),
		},
		Sms: SmsConfig{
			AccountSid: os.Getenv("TWILIO_ACCOUNT_SID"),
			AuthToken:  os.Getenv("TWILIO_AUTH_TOKEN"),
		},
	}
}
