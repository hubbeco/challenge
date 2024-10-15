package tests

import (
	"challenge/models"
	"challenge/services"
	"fmt"
	"testing"
)

func init() {
	// Ativa o uso dos mocks para os envios de email, garantindo que os testes não usem APIs reais
	services.SetEmailMocksEnabled(true)
}

func TestSendEmail_Success(t *testing.T) {
	// Mocka o envio de email, simulando um sucesso no envio
	services.MockSendEmail = func(emailData models.EmailData) error {
		return nil
	}

	// Define os dados de email que serão testados
	emailData := models.EmailData{
		To:      []string{"test@example.com"},
		Subject: "Test Subject",
		Body:    "Test Body",
	}

	// Chama a função SendEmail e verifica se não houve erro
	err := services.SendEmail(emailData)
	if err != nil {
		t.Errorf("Esperava-se que o envio de email fosse bem-sucedido, mas ocorreu um erro: %v", err)
	}
}

func TestSendEmail_Fail(t *testing.T) {
	// Mocka o envio de email, simulando uma falha no envio
	services.MockSendEmail = func(emailData models.EmailData) error {
		return fmt.Errorf("Falha no envio de email")
	}

	// Define os dados de email que serão testados
	emailData := models.EmailData{
		To:      []string{"test@example.com"},
		Subject: "Test Subject",
		Body:    "Test Body",
	}

	// Chama a função SendEmail e verifica se a falha foi corretamente tratada
	err := services.SendEmail(emailData)
	if err == nil || err.Error() != "Falha no envio de email" {
		t.Error("Esperava-se que o envio de email falhasse")
	}
}
