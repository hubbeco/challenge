package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8080"
	}
	return port
}

func GetRecaptchaKey() string {
	return os.Getenv("RECAPTCHA_KEY")
}

func GetRecaptchaURL() string {
	return os.Getenv("RECAPTCHA_URL")
}
