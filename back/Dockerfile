# Etapa de build
FROM golang:1.23.1-alpine AS builder

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
FROM alpine:latest

# Definir diretório de trabalho
WORKDIR /app

# Copiar o binário da etapa anterior
COPY --from=builder /app/myapp .

# Expor a porta que a aplicação usa
EXPOSE 2026

# Comando para rodar a aplicação
CMD ["./myapp"]