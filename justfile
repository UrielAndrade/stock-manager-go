postgres:
    docker run --name postgres-go -e POSTGRES_PASSWORD=123456 -e POSTGRES_DB=estoque -p 5432:5432 -d postgres

up: postgres
    @echo "Aguardando o Postgres iniciar..."
    @sleep 5
    go run main.go

down:
    docker rm -f postgres-go || true

logs:
    docker logs -f postgres-go

ps:
    docker ps -a --filter name=postgres-go