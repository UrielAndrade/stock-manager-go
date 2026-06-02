# Variáveis
BIN_NAME := "estoque-api"

# Inicia o banco de dados e a API via Docker Compose em background (modo detached)
up: down
    @echo "Subindo o ambiente virtual (API e Banco)..."
    docker compose up -d --build

# Inicia o projeto em modo de desenvolvimento atachado aos logs
dev: down
    @echo "Iniciando modo de desenvolvimento com Docker Compose..."
    docker compose up --build

# Para todos os containers do ambiente
down:
    @echo "Parando o ambiente..."
    docker compose down

# Otimiza e compila o binário final (localmente na pasta bin)
build:
    @echo "Compilando binário otimizado na pasta bin/..."
    go build -ldflags="-s -w" -o bin/{{BIN_NAME}} .

# Instala o binário na máquina do usuário localmente (go install)
install: build
    @echo "Instalando a aplicação com go install..."
    go install .

# Apaga o banco de dados, containers e volumes do docker compose
clean-db:
    @echo "Apagando volumes do banco de dados..."
    docker compose down -v

# Ver logs do ambiente
logs:
    docker compose logs -f

# Reseta o banco de dados e reinicia os containers
reset: clean-db up
    @echo "Ambiente resetado!"

# Limpa e formata o projeto localmente
tidy:
    go fmt ./...
    go mod tidy