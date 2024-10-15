package handlers

import (
	"challenge/models"
	"challenge/services"
	"challenge/utils"
	"encoding/json"
	"io"
	"net/http"
)

// HandleContactForm godoc
// @Summary Recebe e processa um formulário de contato
// @Description Este endpoint recebe dados do formulário de contato, valida e envia emails.
// @Tags Contato
// @Accept json
// @Produce json
// @Param contact body models.ContactForm true "Formulário de Contato"
// @Success 201 {201} string "Formulário validado com sucesso, sem retorno no body mas codigo 201"
// @Failure 400 {object} utils.ErrorResponse "Erro de requisição. Possíveis detalhes: 'Name is empty', 'Comment is empty', 'Invalid email format'"
// @Failure 401 {object} utils.ErrorResponse "Captcha incorreto. Detalhe: 'The captcha is incorrect!'"
// @Failure 405 {object} utils.ErrorResponse "Método não permitido. Detalhe: 'Method not allowed'"
// @Failure 415 {object} utils.ErrorResponse "Content-Type incorreto. Detalhe: 'Content-Type must be application/json'"
// @Router /contact [post]
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
	w.Write([]byte("Formulário validado com sucesso! Os emails serão enviados em breve."))
}

func handleRecover(w http.ResponseWriter) {
	if r := recover(); r != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "InternalServerError", "An unexpected error occurred", "/contact")
	}
}
