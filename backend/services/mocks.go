package services

import "challenge/models"

var MockValidateRecaptcha func(string) bool
var MockSendEmail func(data models.EmailData) error

func ValidateRecaptchaMock(response string) bool {
	if MockValidateRecaptcha != nil {
		return MockValidateRecaptcha(response)
	}
	return true
}

func SendEmailMock(emailData models.EmailData) error {
	if MockSendEmail != nil {
		return MockSendEmail(emailData)
	}
	return nil
}
