# Etapa de build
FROM golang:1.23 AS builder
WORKDIR /app

# Copiar dependências e baixar módulos
COPY go.mod go.sum ./
RUN go mod download

# Copiar o código-fonte
COPY . .

# Rodar testes unitários durante o build
# Caso algum teste falhe, o build será interrompido
RUN go test ./... -v

# Compilar a aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -o app

# Etapa final
FROM alpine:latest

# Definir UID e GID como build arguments
ARG USER_ID=1000
ARG GROUP_ID=1000

# Criar grupo e usuário não-root
RUN addgroup -g ${GROUP_ID} appgroup && \
    adduser -D -u ${USER_ID} -G appgroup appuser

WORKDIR /app

# Copiar binário e arquivos necessários para o container
COPY --from=builder /app/app /app/app
COPY ./docs /app/docs
COPY ./tests /app/tests

# Ajustar permissões de arquivos e diretórios
RUN chown -R appuser:appgroup /app

# Definir o usuário não-root para a execução
USER appuser

# Definir variáveis de ambiente
ENV PORT=${PORT}
ENV ORIGINS=${ORIGINS}
ENV RECAPTCHA_KEY=${RECAPTCHA_KEY}
ENV RECAPTCHA_SITE_KEY=${RECAPTCHA_SITE_KEY}
ENV RECAPTCHA_URL=${RECAPTCHA_URL}
ENV MAIL_HOST=${MAIL_HOST}
ENV MAIL_PORT=${MAIL_PORT}
ENV MAIL_SECURE=${MAIL_SECURE}
ENV MAIL_AUTH_USER=${MAIL_AUTH_USER}
ENV MAIL_AUTH_PASS=${MAIL_AUTH_PASS}
ENV TEXT_MAIL_TITLE=${TEXT_MAIL_TITLE}
ENV TEXT_MAIL_BODY=${TEXT_MAIL_BODY}
ENV LOAD_TEST_MODE=${LOAD_TEST_MODE}

# Comando padrão para rodar a aplicação e ajustar permissões após a execução
CMD ["/bin/sh", "-c", "/app/app && chmod 664 /app/tests/result.html"]
