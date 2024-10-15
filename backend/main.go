package main

import (
	"challenge/config"
	_ "challenge/docs"
	"challenge/handlers"
	"challenge/services"
	"fmt"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"html/template"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"
)

var recaptchaSiteKey string

func init() {
	// Carregar a chave do reCAPTCHA no init e definir variáveis de mock para testes de carga
	recaptchaSiteKey = os.Getenv("RECAPTCHA_SITE_KEY")
	if recaptchaSiteKey == "" {
		log.Printf("AVISO: RECAPTCHA_SITE_KEY não configurado. Continuando para cenários de teste...")
	}

	if os.Getenv("LOAD_TEST_MODE") == "true" {
		services.SetRecaptchaMocksEnabled(true)
		services.SetEmailMocksEnabled(true)
	}
	log.Printf("Chave do reCAPTCHA carregada no init: %s", recaptchaSiteKey)
}

// @title Contact Form API
// @version 1.0
// @description API para receber e processar formulários de contato.
// @termsOfService http://swagger.io/terms/

// @contact.name Suporte API
// @contact.url http://www.swagger.io/support
// @contact.email suporte@swagger.io

// @license.name MIT License
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
func main() {

	// Configura a rota do Swagger
	// @Summary Exibir documentação Swagger
	// @Description Exibe a interface gráfica do Swagger para explorar a API.
	// @Tags Swagger
	// @Produce html
	// @Success 200 {string} string "Swagger UI carregado com sucesso"
	// @Router /swagger/ [get]
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	// Logar o número de goroutines a cada 20 segundos
	go func() {
		for {
			numGoroutines := runtime.NumGoroutine()
			log.Printf("Número de Goroutines ativas: %d\n", numGoroutines)
			time.Sleep(20 * time.Second) // Log a cada 5 segundos
		}
	}()

	config.LoadConfig()

	port := config.GetPort()
	origins := config.GetOrigins()

	allowedOrigins := strings.Split(origins, ",")

	// Rota para o formulário de contato com reCAPTCHA dinâmico
	http.HandleFunc("/", serveHTML)

	// @Summary Receber formulário de contato
	// @Description Recebe dados de um formulário de contato, valida e envia emails para o usuário e a empresa.
	// @Tags Contato
	// @Accept json
	// @Produce json
	// @Param contact body models.ContactForm true "Formulário de Contato"
	// @Success 201 {string} string "Formulário validado com sucesso!"
	// @Failure 400 {object} utils.ErrorResponse "Erro de requisição"
	// @Failure 401 {object} utils.ErrorResponse "Captcha incorreto"
	// @Failure 405 {object} utils.ErrorResponse "Método não permitido"
	// @Failure 415 {object} utils.ErrorResponse "Content-Type incorreto"
	// @Router /contact [post]
	http.HandleFunc("/contact", handlers.HandleContactForm)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowCredentials: true,
	})

	handler := corsHandler.Handler(http.DefaultServeMux)

	log.Printf("Servidor rodando na porta %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), handler))
}

// serveHTML renderiza o HTML diretamente no Go com a chave reCAPTCHA injetada
func serveHTML(w http.ResponseWriter, r *http.Request) {
	// HTML template como string
	htmlContent := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Contact Form</title>
		<script src="https://www.google.com/recaptcha/api.js" async defer></script>
	</head>
	<body>
	<h1>Contact Form</h1>
	<form id="contact-form">
		<input type="text" id="name" name="name" placeholder="Your Name"><br>
		<input type="text" id="email" name="email" placeholder="Your Email"><br>
		<textarea id="comment" name="comment" placeholder="Your Comment"></textarea><br>
		<div class="g-recaptcha" data-sitekey="{{.RecaptchaSiteKey}}" data-callback="onSubmit"></div><br>
		<button type="submit">Submit</button>
	</form>

	<div id="response-message" style="margin-top: 20px; color: red; white-space: pre-wrap;"></div>

	<script>
		document.getElementById("contact-form").addEventListener("submit", function(event) {
			event.preventDefault();
			const recaptchaResponse = grecaptcha.getResponse();
			if (!recaptchaResponse) {
				alert("Please complete the reCAPTCHA.");
				return;
			}

			const name = document.getElementById("name").value;
			const email = document.getElementById("email").value;
			const comment = document.getElementById("comment").value;

			const formData = {
				"g-recaptcha-response": recaptchaResponse,
				"name": name,
				"mail": email,
				"comment": comment
			};

			const jsonData = JSON.stringify(formData);

			fetch("/contact", {
				method: "POST",
				headers: {
					"Content-Type": "application/json"
				},
				body: jsonData
			})
			.then(response => {
				if (response.status === 201) {
					document.getElementById("response-message").style.color = "green";
					document.getElementById("response-message").textContent = "Formulário enviado com sucesso!";
				} else {
					return response.json().then(data => {
						const errorResponse = {
							type: data.type,
							title: data.title,
							detail: data.detail,
							instance: data.instance
						};
						document.getElementById("response-message").style.color = "red";
						document.getElementById("response-message").textContent = JSON.stringify(errorResponse, null, 2);
					});
				}
			})
			.catch(error => {
				console.error("Error:", error);
				document.getElementById("response-message").textContent = "Erro ao enviar o formulário.";
			});
		});
	</script>
	</body>
	</html>
	`

	// Printar o valor da chave no log do servidor para verificar
	log.Printf("Chave do reCAPTCHA usada no template: %s", recaptchaSiteKey)

	// Criar o template a partir da string
	tmpl := template.New("captcha")

	// Parsear a string como template
	tmpl, err := tmpl.Parse(htmlContent)
	if err != nil {
		log.Printf("Erro ao criar o template: %v", err)
	}

	// Injetar a variável no template
	err = tmpl.Execute(w, map[string]string{
		"RecaptchaSiteKey": recaptchaSiteKey,
	})
	if err != nil {
		log.Printf("Erro ao renderizar o template: %v", err)
	}
}
