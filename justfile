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

down:
    @if [ "$PODMAN" = "true" ]; then \
        podman rm -f postgres-go || true; \
    else \
        docker rm -f postgres-go || true; \
    fi

logs:
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