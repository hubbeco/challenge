package tests

import (
	"challenge/services"
	"testing"
)

// Inicializa o mock do reCAPTCHA para os testes
func init() {
	services.SetRecaptchaMocksEnabled(true)
}

func TestValidateRecaptcha_Success(t *testing.T) {
	services.MockValidateRecaptcha = func(response string) bool {
		return true
	}

	valid := services.ValidateRecaptcha("valid_captcha")
	if !valid {
		t.Error("Esperava-se que a validação do reCAPTCHA fosse bem-sucedida")
	}
}

func TestValidateRecaptcha_Fail(t *testing.T) {
	services.MockValidateRecaptcha = func(response string) bool {
		return false
	}

	valid := services.ValidateRecaptcha("invalid_captcha")
	if valid {
		t.Error("Esperava-se que a validação do reCAPTCHA falhasse")
	}
}
