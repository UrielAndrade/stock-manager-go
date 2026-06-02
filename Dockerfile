FROM golang:1.25.9-alpine AS builder

WORKDIR /app

# Instalar dependências necessárias para o build
RUN apk add --no-cache git

# Copiar os arquivos de módulos e baixar dependências
COPY go.mod go.sum ./
RUN go mod download

# Copiar o restante do código
COPY . .

# Compilar a aplicação com otimizações (-s -w)
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/estoque-api .

# Estágio final: imagem super leve e limpa
FROM alpine:latest

WORKDIR /app

# O Alpine exige o ca-certificates para chamadas HTTPS externas (caso a API faça requisições)
RUN apk --no-cache add ca-certificates tzdata

# Copiar o binário compilado do estágio anterior
COPY --from=builder /app/estoque-api .

# Copiar o arquivo .env (opcional, pois o docker-compose passará as variáveis, mas útil como fallback)
COPY .env .

EXPOSE 8080

CMD ["./estoque-api"]
