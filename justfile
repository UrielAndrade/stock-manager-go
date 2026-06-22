<<<<<<< HEAD
postgres:
    # Use Podman if PODMAN env var is set, otherwise fallback to Docker
    @if [ "$PODMAN" = "true" ]; then \
        podman run --name postgres-go -e POSTGRES_PASSWORD=123456 -e POSTGRES_DB=estoque -p 5432:5432 -d postgres; \
    else \
        docker run --name postgres-go -e POSTGRES_PASSWORD=123456 -e POSTGRES_DB=estoque -p 5432:5432 -d postgres; \
    fi

# Plug‑and‑play SQLite target for Windows users (no container required)
up-sqlite:
    @echo "Iniciando aplicação com SQLite (plug‑and‑play)"
    DB_TYPE=sqlite go run main.go

# Default `up` now detects environment variable DB_TYPE.
# If DB_TYPE=sqlite, it will use the SQLite target; otherwise it falls back to PostgreSQL.
up: up-sqlite
    @echo "Aguardando o Postgres iniciar..."
    @sleep 5
    go run main.go
=======
# Variáveis
BIN_NAME := "estoque-api"

# Inicia o banco de dados e a API via Docker Compose em background (modo detached)
up: down
    @echo "Subindo o ambiente virtual (API e Banco)..."
    docker compose up -d --build
>>>>>>> 4c8300d5b2fdc3db10fad7ea550d3dec88f36a86

# Inicia o projeto em modo de desenvolvimento atachado aos logs
dev: down
    @echo "Iniciando modo de desenvolvimento com Docker Compose..."
    docker compose up --build

# Para todos os containers do ambiente
down:
<<<<<<< HEAD
    @if [ "$PODMAN" = "true" ]; then \
        podman rm -f postgres-go || true; \
    else \
        docker rm -f postgres-go || true; \
    fi
=======
    @echo "Parando o ambiente..."
    docker compose down
>>>>>>> 4c8300d5b2fdc3db10fad7ea550d3dec88f36a86

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
<<<<<<< HEAD
    @if [ "$PODMAN" = "true" ]; then \
        podman logs -f postgres-go; \
    else \
        docker logs -f postgres-go; \
    fi

reset:
    @if [ "$PODMAN" = "true" ]; then \
        podman rm -f -v postgres-go || true; \
        podman run --name postgres-go -e POSTGRES_PASSWORD=123456 -e POSTGRES_DB=estoque -p 5432:5432 -d postgres; \
    else \
        docker rm -f -v postgres-go || true; \
        docker run --name postgres-go -e POSTGRES_PASSWORD=123456 -e POSTGRES_DB=estoque -p 5432:5432 -d postgres; \
    fi
    @echo "Banco resetado e todos os dados foram apagados."
    just up
    @sleep 5
    go run main.go

ps:
    @if [ "$PODMAN" = "true" ]; then \
        podman ps -a --filter name=postgres-go; \
    else \
        docker ps -a --filter name=postgres-go; \
    fi
=======
    docker compose logs -f

# Reseta o banco de dados e reinicia os containers
reset: clean-db up
    @echo "Ambiente resetado!"

# Limpa e formata o projeto localmente
tidy:
    go fmt ./...
    go mod tidy
>>>>>>> 4c8300d5b2fdc3db10fad7ea550d3dec88f36a86
