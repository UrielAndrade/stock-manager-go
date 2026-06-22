# Podman / Docker conditional commands

# Start PostgreSQL container (Podman if PODMAN=true, else Docker)
postgres:
	@if [ "$PODMAN" = "true" ]; then \
		podman run --name postgres-go -e POSTGRES_PASSWORD=123456 -e POSTGRES_DB=estoque -p 5432:5432 -d postgres; \
	else \
		docker run --name postgres-go -e POSTGRES_PASSWORD=123456 -e POSTGRES_DB=estoque -p 5432:5432 -d postgres; \
	fi

# SQLite target (no container)
up-sqlite:
	@echo "Iniciando aplicação com SQLite (plug‑and‑play)"
	DB_TYPE=sqlite go run main.go

# Default up target: use SQLite if DB_TYPE=sqlite, otherwise start PostgreSQL then run API
up:
	@if [ "$$DB_TYPE" = "sqlite" ]; then \
		$(MAKE) up-sqlite; \
	else \
		$(MAKE) postgres; \
		@echo "Aguardando o Postgres iniciar..."; \
		@sleep 5; \
		go run main.go; \
	fi

# Stop and remove PostgreSQL container
down:
	@if [ "$PODMAN" = "true" ]; then \
		podman rm -f postgres-go || true; \
	else \
		docker rm -f postgres-go || true; \
	fi

# Stream logs from PostgreSQL container
logs:
	@if [ "$PODMAN" = "true" ]; then \
		podman logs -f postgres-go; \
	else \
		docker logs -f postgres-go; \
	fi

# Reset database container and restart application
reset:
	@if [ "$PODMAN" = "true" ]; then \
		podman rm -f -v postgres-go || true; \
		podman run --name postgres-go -e POSTGRES_PASSWORD=123456 -e POSTGRES_DB=estoque -p 5432:5432 -d postgres; \
	else \
		docker rm -f -v postgres-go || true; \
		docker run --name postgres-go -e POSTGRES_PASSWORD=123456 -e POSTGRES_DB=estoque -p 5432:5432 -d postgres; \
	fi
	@echo "Banco resetado e todos os dados foram apagados."
	$(MAKE) up
	@sleep 5
	go run main.go

# List containers related to the project
ps:
	@if [ "$PODMAN" = "true" ]; then \
		podman ps -a --filter name=postgres-go; \
	else \
		docker ps -a --filter name=postgres-go; \
	fi

# Build and install
BIN_NAME := "estoque-api"

build:
	@echo "Compilando binário otimizado na pasta bin/..."
	go build -ldflags="-s -w" -o bin/$(BIN_NAME) .

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
