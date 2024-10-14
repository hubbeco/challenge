package services

import (
	"challenge/models"
	"time"
)

// Mocks são usados para testes, evitando o uso de APIs de terceiros, que podem ser pagas ou limitar o número de requisições.
// Os delays simulam o tempo de resposta das APIs externas para garantir que o comportamento da aplicação seja testado de forma realística.

var MockValidateRecaptcha func(string) bool
var MockSendEmail func(data models.EmailData) error

// Simula a validação do reCAPTCHA com um delay para imitar a latência da API real.
func ValidateRecaptchaMock(response string) bool {
	if MockValidateRecaptcha != nil {
		return MockValidateRecaptcha(response)
	}
	time.Sleep(250 * time.Millisecond) // Simula o tempo de resposta da API real
	return true
}

// Simula o envio de email com um delay para imitar o tempo de envio real.
func SendEmailMock(emailData models.EmailData) error {
	if MockSendEmail != nil {
		return MockSendEmail(emailData)
	}
	time.Sleep(3 * time.Second) // Simula o tempo de envio do email
	return nil
}
