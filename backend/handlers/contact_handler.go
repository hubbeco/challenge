package handlers

import (
	"challenge/models"
	"challenge/services"
	"challenge/utils"
	"encoding/json"
	"io"
	"net/http"
	"regexp"
)

func HandleContactForm(w http.ResponseWriter, r *http.Request) {
	defer handleRecover(w)

	if r.Method != http.MethodPost {
		utils.RespondWithError(w, http.StatusMethodNotAllowed, "MethodNotAllowed", "Method not allowed", "/contact")
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		utils.RespondWithError(w, http.StatusUnsupportedMediaType, "UnsupportedMediaType", "Content-Type must be application/json", "/contact")
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "BadRequestError", "Invalid request body", "/contact")
		return
	}

	var form models.ContactForm
	err = json.Unmarshal(body, &form)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "BadRequestError", "Invalid JSON format", "/contact")
		return
	}

	if form.Name == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "BadRequestError", "The name is empty", "/contact")
		return
	}

	if form.Comment == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "BadRequestError", "The comment is empty", "/contact")
		return
	}

	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	if !emailRegex.MatchString(form.Email) {
		utils.RespondWithError(w, http.StatusBadRequest, "BadRequestError", "The email is invalid", "/contact")
		return
	}

	if !services.ValidateRecaptcha(form.RecaptchaResponse) {
		utils.RespondWithError(w, http.StatusUnauthorized, "UnauthorizedError", "The captcha is incorrect!", "/contact")
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func handleRecover(w http.ResponseWriter) {
	if r := recover(); r != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "InternalServerError", "An unexpected error occurred", "/contact")
	}
}
