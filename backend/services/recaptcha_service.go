package services

import (
	"challenge/config"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"
)

var useRecaptchaMocks = false

// SetRecaptchaMocksEnabled ativa ou desativa o uso de mocks para o reCAPTCHA.
// Quando habilitado, evita chamadas reais à API externa do recaptcha e usa o mock para simular o comportamento.
func SetRecaptchaMocksEnabled(enabled bool) {
	useRecaptchaMocks = enabled
}

// ValidateRecaptcha verifica a resposta do reCAPTCHA.
// Se o modo de mock estiver habilitado, chama o mock para evitar o uso de API externa.
// Caso contrário, faz uma requisição a real para à API do reCAPTCHA e valida a resposta.
func ValidateRecaptcha(recaptchaResponse string) bool {

	// Timeout de 500ms para evitar que caso a requisicao seja perdida os recursos sejam liberados
	client := &http.Client{
		Timeout: 500 * time.Millisecond,
	}

	// Se os mocks estiverem habilitados, usa o mock para validar o recaptcha.
	if useRecaptchaMocks {
		return ValidateRecaptchaMock(recaptchaResponse)
	}

	// Configura a URL e os dados necessários para a requisição à API do reCAPTCHA.
	secretKey := config.GetRecaptchaKey()
	recaptchaURL := config.GetRecaptchaURL()

	// Define os parâmetros para a requisição.
	data := url.Values{}
	data.Set("secret", secretKey)
	data.Set("response", recaptchaResponse)

	// Envia uma requisição POST para validar o reCAPTCHA.
	resp, err := client.PostForm(recaptchaURL, data)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	// Lê a resposta do corpo da requisição.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false
	}

	// Estrutura para armazenar o resultado da validação do reCAPTCHA.
	var result struct {
		Success bool `json:"success"`
	}

	// Faz o parse do JSON retornado pela API do reCAPTCHA.
	err = json.Unmarshal(body, &result)
	return err == nil && result.Success
}
