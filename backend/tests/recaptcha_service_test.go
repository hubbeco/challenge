package tests

import (
	"challenge/services"
	"testing"
)

// Ativa o mock do recaptcha
func init() {
	services.SetRecaptchaMocksEnabled(true)
}

// Testa recaptcha válido
func TestValidateRecaptcha_Success(t *testing.T) {
	services.MockValidateRecaptcha = func(response string) bool {
		return true
	}

	if !services.ValidateRecaptcha("valid_captcha") {
		t.Error("Recaptcha deveria ser válido")
	}
}

// Testa recaptcha inválido
func TestValidateRecaptcha_Fail(t *testing.T) {
	services.MockValidateRecaptcha = func(response string) bool {
		return false
	}

	if services.ValidateRecaptcha("invalid_captcha") {
		t.Error("Recaptcha deveria ser inválido")
	}
}
