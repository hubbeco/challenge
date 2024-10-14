package config

import (
	"github.com/joho/godotenv"
	"os"
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
	}
}

func GetOrigins() string {
	origins := os.Getenv("ORIGINS")
	if origins == "" {
		return ""
	}
	return origins
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

func GetSMTPHost() string {
	return os.Getenv("MAIL_HOST")
}

func GetSMTPPort() string {
	return os.Getenv("MAIL_PORT")
}

func GetSMTPAuthUser() string {
	return os.Getenv("MAIL_AUTH_USER")
}

func GetSMTPAuthPass() string {
	return os.Getenv("MAIL_AUTH_PASS")
}

func GetMailTitle() string {
	return os.Getenv("TEXT_MAIL_TITLE")
}

func GetMailBody() string {
	return os.Getenv("TEXT_MAIL_BODY")
}
