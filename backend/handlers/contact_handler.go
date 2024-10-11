package handlers

import (
	"challenge/models"
	"challenge/services"
	"challenge/utils"
	"encoding/json"
	"io"
	"net/http"
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

	err = services.ProcessContactForm(form)
	if err != nil {
		if err.Error() == "UnauthorizedError: The captcha is incorrect!" {
			utils.RespondWithError(w, http.StatusUnauthorized, "UnauthorizedError", "The captcha is incorrect!", "/contact")
		} else {
			utils.RespondWithError(w, http.StatusBadRequest, "BadRequestError", err.Error(), "/contact")
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func handleRecover(w http.ResponseWriter) {
	if r := recover(); r != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "InternalServerError", "An unexpected error occurred", "/contact")
	}
}
