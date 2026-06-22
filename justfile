# Podman / Docker conditional commands

# Start PostgreSQL container (Podman if PODMAN=true, else Docker)
postgres:
	@if [ "${PODMAN:-}" = "true" ]; then \
		if [ -z "$(podman ps -a -q --filter name=^postgres-go$)" ]; then \
			podman run --name postgres-go -e POSTGRES_PASSWORD=123456 -e POSTGRES_DB=estoque -p 5432:5432 -d postgres; \
		elif [ -z "$(podman ps -q --filter name=^postgres-go$)" ]; then \
			podman start postgres-go; \
		fi; \
	else \
		if [ -z "$(docker ps -a -q --filter name=^postgres-go$)" ]; then \
			docker run --name postgres-go -e POSTGRES_PASSWORD=123456 -e POSTGRES_DB=estoque -p 5432:5432 -d postgres; \
		elif [ -z "$(docker ps -q --filter name=^postgres-go$)" ]; then \
			docker start postgres-go; \
		fi; \
	fi

# SQLite target (no container)
up-sqlite:
	@echo "Iniciando aplicação com SQLite (plug‑and‑play)"
	DB_TYPE=sqlite go run main.go

# Default up target: use SQLite if DB_TYPE=sqlite, otherwise start PostgreSQL then run API
up:
	@if [ "${DB_TYPE:-}" = "sqlite" ]; then \
		just up-sqlite; \
	else \
		echo "Iniciando aplicação completa no Docker..."; \
		docker compose up --build; \
	fi

# Stop and remove PostgreSQL container
down:
	@if [ "${PODMAN:-}" = "true" ]; then \
		podman rm -f postgres-go || true; \
	else \
		docker compose down; \
		docker rm -f postgres-go || true; \
	fi

# Stream logs from PostgreSQL container
logs:
	@if [ "${PODMAN:-}" = "true" ]; then \
		podman logs -f postgres-go; \
	else \
		docker compose logs -f; \
	fi

# Reset database container and restart application
reset:
	@if [ "${PODMAN:-}" = "true" ]; then \
		podman rm -f -v postgres-go || true; \
	else \
		docker compose down -v; \
	fi
	@echo "Banco resetado e todos os dados foram apagados."
	just up

# List containers related to the project
ps:
	@if [ "${PODMAN:-}" = "true" ]; then \
		podman ps -a --filter name=postgres-go; \
	else \
		docker compose ps -a; \
	fi

# Build and install
BIN_NAME := "estoque-api"

build:
	@echo "Compilando binário otimizado na pasta bin/..."
	go build -ldflags="-s -w" -o bin/{{BIN_NAME}} .

install: build
	@echo "Instalando a aplicação com go install..."
	go install .

# Remove Docker Compose volumes (if using compose elsewhere)
clean-db:
	@echo "Apagando volumes do banco de dados..."
	docker compose down -v

# Format and tidy Go modules
tidy:
	go fmt ./...
	go mod tidy
