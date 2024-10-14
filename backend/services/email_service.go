package services

import (
	"challenge/config"
	"challenge/models"
	"context"
	"fmt"
	"net/smtp"
	"regexp"
	"strings"
	"sync"
	"time"
)

var useEmailMocks = false

// SetEmailMocksEnabled ativa ou desativa o uso de mocks para o envio de emails durante os testes
// Afim de não gastar recursos em uma api de terceiros
func SetEmailMocksEnabled(enabled bool) {
	useEmailMocks = enabled
}

// SendEmail envia um email para o(s) destinatário(s) especificado(s).
// Utiliza o protocolo SMTP com as configurações lidas das variáveis de ambiente, testado com o servico mailtrap.
// Em caso de falha, tenta reenviar o email dentro de um timeout de 7 segundos, após, libera o recurso alocado.
func SendEmail(emailData models.EmailData) error {

	// Simula o envio de email para cenários de teste, configurado delay de 3 segundos.
	// 3 segundos é o delay real para codigo 201 quando depende apenas do email,o captcha custa 250ms para responder
	if useEmailMocks {
		return SendEmailMock(emailData)
	}

	// Configurações do servidor SMTP
	smtpHost := config.GetSMTPHost()
	smtpPort := config.GetSMTPPort()
	authUser := config.GetSMTPAuthUser()
	authPass := config.GetSMTPAuthPass()

	auth := smtp.PlainAuth("", authUser, authPass, smtpHost)

	// Construção da mensagem de email
	message := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s",
		authUser,
		strings.Join(emailData.To, ","),
		emailData.Subject,
		emailData.Body)

	smtpAddress := fmt.Sprintf("%s:%s", smtpHost, smtpPort)

	// Contexto com timeout de 7 segundos para garantir que o envio de email não bloqueie recursos indefinidamente
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	errChan := make(chan error, 1)

	// Goroutine para envio do email sem bloquear o fluxo principal
	go func() {
		err := smtp.SendMail(
			smtpAddress,
			auth,
			authUser,
			emailData.To,
			[]byte(message),
		)
		errChan <- err
	}()

	// Tratamento de timeout ou erro no envio de email
	select {
	case <-ctx.Done():
		return fmt.Errorf("timeout: email sending took too long")
	case err := <-errChan:
		if err != nil {
			return fmt.Errorf("erro ao enviar email: %v", err)
		}
	}

	return nil
}

// ProcessContactForm processa os dados do formulário de contato.
// Realiza a validação do formulário, verifica o reCAPTCHA e envia emails ao usuário e à empresa.
// O envio de emails é feito de forma assíncrona para tornar a API mais responsiva.
func ProcessContactForm(form models.ContactForm) error {

	// Valida os dados do formulário
	if err := ValidateForm(form); err != nil {
		return fmt.Errorf(err.Error())
	}

	// Valida o recaptcha, caso falhe, retorna erro de autorização com detalhes.
	if !ValidateRecaptcha(form.RecaptchaResponse) {
		return fmt.Errorf("UnauthorizedError: The captcha is incorrect!")
	}

	// Goroutine para realizar o envio dos emails de forma assíncrona
	// Usamos goroutines para não bloquear a resposta da API enquanto os emails são enviados
	go func() {
		var wg sync.WaitGroup
		errChan := make(chan error, 2)

		// Usamos o WaitGroup para controlar o sincronismo das goroutines que enviam os emails
		wg.Add(2)

		// Envio de email para o usuário
		go func() {
			defer wg.Done()
			if err := SendUserEmail(form); err != nil {
				errChan <- err
			}
		}()

		// Envio de email para a empresa
		go func() {
			defer wg.Done()
			if err := SendCompanyEmail(form); err != nil {
				errChan <- err
			}
		}()

		// Aguarda a conclusão das duas goroutines de envio de email
		wg.Wait()
		close(errChan)

		// Verifica se houve erro no envio de algum dos emails
		var emailErr error
		for err := range errChan {
			if emailErr == nil {
				emailErr = err
			}
		}

		// Aqui pode ser implementada uma lógica para alertar em caso de falha de envio como um log por exemplo
		if emailErr != nil {
			// Exemplo: logar o erro ou realizar uma nova tentativa de envio
		} else {

		}
	}()

	return nil
}

// ValidateForm realiza a validação dos campos do formulário de contato.
// O ideal é que as validacoes de name e comment sejam executadas antes do email, pois verificam somente string vazia
// Já o email exige a validação de um regex que gasta mais recursos em que pese
func ValidateForm(form models.ContactForm) error {

	if form.Name == "" {
		return fmt.Errorf("Name is empty")
	}
	if form.Comment == "" {
		return fmt.Errorf("Comment is empty")
	}

	// Verifica se o formato do email é válido
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	if !emailRegex.MatchString(form.Email) {
		return fmt.Errorf("Invalid email format")
	}

	return nil
}

// SendUserEmail prepara e envia o email de confirmação para o usuário.
func SendUserEmail(form models.ContactForm) error {
	userBody := strings.Replace(config.GetMailBody(), "{name}", form.Name, -1)
	userBody = strings.Replace(userBody, "{email}", form.Email, -1)
	userBody = strings.Replace(userBody, "{comment}", form.Comment, -1)

	userEmailData := models.EmailData{
		To:      []string{form.Email},
		Subject: config.GetMailTitle(),
		Body:    userBody,
	}

	return SendEmail(userEmailData)
}

// SendCompanyEmail prepara e envia o email para a empresa com os detalhes do contato recebido.
func SendCompanyEmail(form models.ContactForm) error {
	companyBody := fmt.Sprintf("Nome: %s\nEmail: %s\nComentário: %s", form.Name, form.Email, form.Comment)

	companyEmailData := models.EmailData{
		To:      []string{"email_da_empresa@example.com"}, // Substitua pelo email real da empresa
		Subject: "Novo contato recebido",
		Body:    companyBody,
	}

	return SendEmail(companyEmailData)
}
