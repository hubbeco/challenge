package tests

import (
	"challenge/models"
	"challenge/services"
	"fmt"
	"testing"
)

// Inicializa o mock de envio de email para os testes
func init() {
	services.SetEmailMocksEnabled(true)
}

func TestSendEmail_Success(t *testing.T) {
	services.MockSendEmail = func(emailData models.EmailData) error {
		return nil
	}

	emailData := models.EmailData{
		To:      []string{"test@example.com"},
		Subject: "Test Subject",
		Body:    "Test Body",
	}

	err := services.SendEmail(emailData)
	if err != nil {
		t.Errorf("Esperava-se que o envio de email fosse bem-sucedido, mas ocorreu um erro: %v", err)
	}
}

func TestSendEmail_Fail(t *testing.T) {
	services.MockSendEmail = func(emailData models.EmailData) error {
		return fmt.Errorf("Falha no envio de email")
	}

	emailData := models.EmailData{
		To:      []string{"test@example.com"},
		Subject: "Test Subject",
		Body:    "Test Body",
	}

	err := services.SendEmail(emailData)
	if err == nil || err.Error() != "Falha no envio de email" {
		t.Error("Esperava-se que o envio de email falhasse")
	}
}
