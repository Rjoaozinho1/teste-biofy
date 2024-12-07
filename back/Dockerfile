# Etapa de build
FROM golang:1.23-alpine AS builder

# Definir diretório de trabalho
WORKDIR /app

# Copiar os arquivos go.mod e go.sum para o contêiner
COPY go.mod go.sum ./

# Baixar as dependências
RUN go mod download

# Copiar o código-fonte para o contêiner
COPY . .

# Compilar o binário
RUN go build -o myapp .

# Etapa final - execução
FROM alpine:3.18

# Criar um usuário não-root para rodar a aplicação
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Definir diretório de trabalho
WORKDIR /app

# Copiar o binário da etapa anterior
COPY --from=builder /app/myapp .

# Alterar para o usuário não-root
USER appuser

# Expor a porta que a aplicação usa
EXPOSE 2026

# Comando para rodar a aplicação
CMD ["./myapp"]