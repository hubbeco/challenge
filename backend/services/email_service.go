package services

import (
	"challenge/config"
	"challenge/models"
	"fmt"
	"net/smtp"
	"regexp"
	"strings"
	"sync"
)

type EmailData struct {
	To      []string
	Subject string
	Body    string
}

func SendEmail(emailData EmailData) error {
	smtpHost := config.GetSMTPHost()
	smtpPort := config.GetSMTPPort()
	authUser := config.GetSMTPAuthUser()
	authPass := config.GetSMTPAuthPass()

	auth := smtp.PlainAuth("", authUser, authPass, smtpHost)

	message := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s",
		authUser,
		strings.Join(emailData.To, ","),
		emailData.Subject,
		emailData.Body)

	smtpAddress := fmt.Sprintf("%s:%s", smtpHost, smtpPort)

	err := smtp.SendMail(
		smtpAddress,
		auth,
		authUser,
		emailData.To,
		[]byte(message),
	)

	if err != nil {
		return fmt.Errorf("Erro ao enviar email: %v", err)
	}
	return nil
}

func ProcessContactForm(form models.ContactForm) error {

	if err := ValidateForm(form); err != nil {
		return fmt.Errorf(err.Error())
	}

	if !ValidateRecaptcha(form.RecaptchaResponse) {
		return fmt.Errorf("UnauthorizedError: The captcha is incorrect!")
	}

	var wg sync.WaitGroup
	errChan := make(chan error, 2)

	wg.Add(2)

	go func() {
		defer wg.Done()
		if err := SendUserEmail(form); err != nil {
			errChan <- err
		}
	}()

	go func() {
		defer wg.Done()
		if err := SendCompanyEmail(form); err != nil {
			errChan <- err
		}
	}()

	wg.Wait()
	close(errChan)

	var emailErr error
	for err := range errChan {
		if emailErr == nil {
			emailErr = err
		}
	}

	if emailErr != nil {
		return fmt.Errorf("Falha ao enviar email: %v", emailErr)
	}

	return nil
}

func ValidateForm(form models.ContactForm) error {
	if form.Name == "" {
		return fmt.Errorf("Name is empty")
	}
	if form.Comment == "" {
		return fmt.Errorf("Comment is empty")
	}

	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	if !emailRegex.MatchString(form.Email) {
		return fmt.Errorf("Invalid email format")
	}

	return nil
}

func SendUserEmail(form models.ContactForm) error {
	userBody := strings.Replace(config.GetMailBody(), "{name}", form.Name, -1)
	userBody = strings.Replace(userBody, "{email}", form.Email, -1)
	userBody = strings.Replace(userBody, "{comment}", form.Comment, -1)

	userEmailData := EmailData{
		To:      []string{form.Email},
		Subject: config.GetMailTitle(),
		Body:    userBody,
	}

	return SendEmail(userEmailData)
}

func SendCompanyEmail(form models.ContactForm) error {
	companyBody := fmt.Sprintf("Nome: %s\nEmail: %s\nComentário: %s", form.Name, form.Email, form.Comment)

	companyEmailData := EmailData{
		To:      []string{"email_da_empresa@example.com"}, // Substitua pelo email real da empresa
		Subject: "Novo contato recebido",
		Body:    companyBody,
	}

	return SendEmail(companyEmailData)
}
