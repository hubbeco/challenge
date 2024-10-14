
# Microserviço de Formulário de Contato

Este projeto é um microserviço para um formulário de contato que envia e-mails tanto para a empresa quanto para o usuário, utilizando validação de reCAPTCHA. A aplicação é configurável dinamicamente via variáveis de ambiente, o que permite fácil integração com diferentes ambientes.

## Funcionalidades

- Validação de reCAPTCHA (v2 ou hCaptcha)
- Envio de e-mails para a empresa e para o usuário
- Configuração flexível via variáveis de ambiente

## Requisitos

- Linux
- Docker

## Configuração

O microserviço é configurado via variáveis de ambiente, que estão definidas no Dockerfile e docker-compose. Essas variáveis recebem valores no momento de execução do container.

Variáveis de ambiente:

As variáveis de ambiente abaixo são usadas para configurar o microserviço de formulário de contato:

- **`PORT`**: Define a porta na qual o serviço será executado.
- Exemplo: `8900`

- **`ORIGINS`**: Define as origens permitidas para CORS (Cross-Origin Resource Sharing).
- Exemplo: `https://example.com`

- **`RECAPTCHA_KEY`**: Chave secreta da API do reCAPTCHA, usada para validar o reCAPTCHA no backend.
- Exemplo: `your-recaptcha-secret-key`

- **`RECAPTCHA_SITE_KEY`**: Chave pública do reCAPTCHA, utilizada no frontend para exibir o widget de verificação.
- Exemplo: `your-recaptcha-site-key`

- **`RECAPTCHA_URL`**: URL usada para verificar a resposta do reCAPTCHA.
- Exemplo: `https://www.google.com/recaptcha/api/siteverify`

- **`MAIL_HOST`**: Host do servidor SMTP para envio de e-mails.
- Exemplo: `smtp.mailtrap.io`

- **`MAIL_PORT`**: Porta do servidor SMTP.
- Exemplo: `587`

- **`MAIL_SECURE`**: Define se o envio de e-mail deve ser feito de forma segura (usando SSL/TLS).
- Exemplo: `false` (para não usar SSL) ou `true` (para usar SSL)

- **`MAIL_AUTH_USER`**: Nome de usuário usado para autenticação no servidor SMTP.
- Exemplo: `your-smtp-username`

- **`MAIL_AUTH_PASS`**: Senha usada para autenticação no servidor SMTP.
- Exemplo: `your-smtp-password`

- **`TEXT_MAIL_TITLE`**: Título do e-mail enviado para o usuário.
- Exemplo: `"Contato recebido com sucesso"`

- **`TEXT_MAIL_BODY`**: Corpo do e-mail enviado ao usuário.
- Exemplo: `"Obrigado por entrar em contato conosco, responderemos em breve."`

- **`LOAD_TEST_MODE`**: Modo de teste de carga, que ativa mocks para reCAPTCHA e e-mails.
- Exemplo: `true` (ativa o modo de teste de carga) ou `false` (desativa o modo de teste)

## Como Executar

### Usando Docker

1. Garanta que o usuário tenha o docker instalado e acesso ao grupo docker:

-Instalando o docker.

```bash
sudo snap install docker
```
ou

```bash
sudo apt  install docker.io      # version 24.0.7-0ubuntu4.1
```
> Nota: Comandos para instalar o docker.


-Criamos o grupo docker e adicionamos o úsuarios com os seguintes comando


```bash
sudo groupadd docker
sudo usermod -aG docker $USER
```
> Nota: Comandos para instalar adicionar o docker group.

2. Reinicie o sistema para aplicar as mudancas.

3. Navegue até o diretório `backend`.

4. Construa a imagem Docker:

```bash
docker build --build-arg USER_ID=$(id -u) --build-arg GROUP_ID=$(id -g) -t my-go-app .
```

5. Execute o container:

```bash
docker run -p 8800:8800 -e PORT=8800 -e LOAD_TEST_MODE=true my-go-app
```

#### Sintaxe alternativa para produção:

```bash
docker run -e PORT=8900 -e RECAPTCHA_SITE_KEY=<Sua chave do site key> -e RECAPTCHA_KEY=<sua chave do backend> -e RECAPTCHA_URL=<"link da api que valida seu recaptcha"> -e LOAD_TEST_MODE=false -e MAIL_HOST=<Api que envia o email> -e MAIL_AUTH_USER=<Login da api de email> -e MAIL_AUTH_PASS=<Senha da api de email> -e MAIL_PORT=465 -e MAIL_SECURE=false -p 8900:8900 my-go-app
```

### Usando Docker Compose (Ideal para Testes de Carga)

1. Siga os passos 1, 2 e 3 da sessão anterior.

2. Construa a imagem Docker:

```bash
docker build --build-arg USER_ID=$(id -u) --build-arg GROUP_ID=$(id -g) -t my-go-app .
```

3. Execute o container.

#### Sintaxe para produção:

```bash
PORT=8900 RECAPTCHA_SITE_KEY=<Sua chave do site key> RECAPTCHA_KEY=<sua chave do backend> RECAPTCHA_URL=<"link da api que valida seu recaptcha"> MAIL_HOST=sandbox.smtp.mailtrap.io MAIL_PORT=465 MAIL_SECURE=false MAIL_AUTH_USER=<Login da api de email> MAIL_AUTH_PASS=<Senha da api de email> LOAD_TEST_MODE=false docker-compose up --build
```
> Nota: A sintaxe alterativa para produção, que ira ser cita ao logo deste, trata de uma sintaxe com todas as variaveis configuradas para um uso real.

#### Sintaxe para testes:
```bash
PORT=8900 LOAD_TEST_MODE=true UID=$(id -u) GID=$(id -g) docker-compose up --build
```

## Testes
### Testes Unitários
Para rodar os testes unitários e automatizados:

- **IDE**: Execute o módulo `tests`.
- **Durante o build**: Execute o seguinte comando:

```bash
docker build --build-arg USER_ID=$(id -u) --build-arg GROUP_ID=$(id -g) -t my-go-app .
```
> Nota: Os resultados dos testes unitários automatizados aparecerão como log no terminal durante o processo de build. Após o build, não será mais possível rodar os testes, pois, para otimizar a imagem para ambientes de produção, o compilador Go é removido, deixando apenas o binário nativo para execução. Para rodar os testes novamente, será necessário rebuildar a imagem.


### Testes com Postman
#### 1. Configurando o endpoint de envio de contato:

- **Método:** `POST`
- **URL:** `http://localhost:<SUA-PORTA>/contact`
#### Cabeçalhos (Headers):
- **Content-Type:** `application/json`
#### Body:
Selecione o tipo de body como `raw` e o formato como `JSON`. Utilize o seguinte exemplo de payload:
```json
  {
  "g-recaptcha-response": "sua resposta do recaptcha",
  "comment": "Comentário de exemplo",
  "name": "Seu nome",
  "mail": "seu.email@example.com"
  }
```
### 2. Executando o teste de envio correto:
- Após configurar o body, clique em "Send" para enviar a requisição.
- **Resposta esperada:** Código de status `201` para envio correto.

> Nota: Caso execute com LOAD_TEST_MODE=true qualquer recaptcha será aceito, caso execute com LOAD_TEST_MODE=false, será preciso configurar chaves de validação do recaptcha para que funcione, do contrário retornará 401 Unauthorized.

### 3. Testando erro de captcha:

- Altere o valor de `"g-recaptcha-response"` para um valor incorreto.
- **Resposta esperada:** Código de status `401` com um JSON de erro conforme a RFC7807:

  ```json
  {
  "type": "about:blank",
  "title": "UnauthorizedError",
  "detail": "The captcha is incorrect!",
  "instance": "/api-endpoint"
  }
  ```

### 4. Testando erro de usuário (email inválido):

- Altere o campo `"mail"` para um email inválido, por exemplo, `"seu.email@"`.
- **Resposta esperada:** Código de status `400` com um JSON de erro conforme a RFC7807:

  ```json
  {
  "type": "about:blank",
  "title": "BadRequestError",
  "detail": "The email is invalid",
  "instance": "/api-endpoint"
  }
    ```

### 5. Testando erro de usuário (nome vazio):

- Deixe o campo `"name"` vazio.
- **Resposta esperada:** Código de status `400` com um JSON de erro conforme a RFC7807:
-
    ```json
    {
    "type": "about:blank",
    "title": "BadRequestError",
    "detail": "The name is empty",
    "instance": "/api-endpoint"
    }
    ```

Essas são as instruções básicas para realizar testes com o Postman em sua API de formulário de contato. Certifique-se de que a API esteja rodando e acessível no `localhost` ou no domínio configurado.

### Testes de Carga

Os testes de carga foram executados com o framework k6 e podem ser ajustados no arquivo `load_test.js`. O docker-compose está configurado para rodar o container com 50% de CPU e 64MB de RAM, o arquivo result.html disponivel no resultado tests representa o desempenho com um i5 1340p em modo performance.

Comando para rodar o teste de carga:

```bash
docker-compose exec --user "$(id -u):$(id -g)" k6 sh -c "umask 002 && k6 run /tests/load_test.js && chmod 664 /tests/result.html"
```

> Nota: O teste de carga foi projetado para rodar via docker-compose, pois ele gerencia tanto o container da API quanto o de teste.


Após a execução dos testes de carga, um arquivo HTML é gerado no diretório `backend/tests`. Esse arquivo contém os resultados dos testes, incluindo informações sobre o desempenho da aplicação sob diferentes condições de carga.

O arquivo gerado é:

- **`result.html`**: Este arquivo exibe um relatório detalhado dos testes de carga executados, incluindo o tempo de resposta das requisições e outras métricas importantes. Ele pode ser acessado no caminho:

```bash
backend/tests/result.html
```
## Documentação da API

### POST /contact

Esse endpoint recebe os dados do formulário, valida o reCAPTCHA e envia e-mails para o usuário e para a empresa.

#### Exemplo de Sucesso (201):

```json
{
"message": "O formulário foi validado e os e-mails serão enviados."
}
```

#### Exemplo de Erros:

- **400 BadRequestError**: Falha na validação dos dados.
- **401 UnauthorizedError**: O reCAPTCHA é inválido ou incorreto.
- **405 MethodNotAllowed**: Apenas o método POST é aceito.
- **415 UnsupportedMediaType**: O tipo de conteúdo deve ser `application/json`.
- **500 InternalServerError**: Erro interno do servidor.

### Documentação Swagger

A documentação completa da API pode ser acessada via Swagger:

- **URL**: `http://localhost:<SUA-PORTA>/swagger/index.html`

## /static/recaptcha_test.html

Quando o container está em execução este é um formulário HTML estático disponível no caminho `http://localhost:<PORTA-DO-SEU-CONTAINER>/static/recaptcha_test.html` através do navegador. Ele foi desenvolvido para testar a integração do reCAPTCHA e o envio de e-mails. Esse formulário é utilizado apenas para testes internos, e a implementação do HTML diretamente no código Go não é uma boa prática de desenvolvimento, mas foi feita neste caso para fins de teste rápido.

## Observações Adicionais

### Go Routines

Durante os testes de carga, a aplicação faz uso intenso de Go Routines, permitindo que o microserviço lide com múltiplas requisições simultaneamente. A cada 15 segundos, o número de Go Routines ativas é registrado através de um log no terminal de execução do container, evidenciando como o Go Runtime gerencia a concorrência e paralelismo de forma eficiente.

### Vídeo de Teste de Carga

Um vídeo demonstrando o teste de carga e a análise de desempenho da aplicação pode ser visualizado neste link:

[Link do vídeo](https://youtu.be/vZqO39JPCk0)
