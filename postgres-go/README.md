# PostgreSQL Docker Setup

This project sets up a PostgreSQL database using Docker. Below are the instructions to build and run the Docker container.

## Prerequisites

- Docker installed on your machine.

## Project Structure

```
postgres-go
├── Dockerfile
├── .env
├── .dockerignore
└── README.md
```

## Environment Configuration

The `.env` file contains the following environment variables for the PostgreSQL database configuration:

- `DB_HOST`: The hostname of the database (default is `localhost`).
- `DB_PORT`: The port on which the database will run (default is `5432`).
- `DB_USER`: The username for the database (default is `postgres`).
- `DB_PASSWORD`: The password for the database user.
- `DB_NAME`: The name of the database to create (default is `estoque`).
- `DB_SSLMODE`: SSL mode for the database connection (default is `disable`).

## Building the Docker Image

To build the Docker image, navigate to the project directory and run the following command:

```
docker build -t postgres-go .
```

## Running the Docker Container

To run the PostgreSQL database container, use the following command:

```
docker run --name postgres-go -e POSTGRES_PASSWORD=123456 -e POSTGRES_DB=estoque -p 5432:5432 -d postgres
```

This command will:

- Create a new container named `postgres-go`.
- Set the PostgreSQL password to `123456`.
- Create a database named `estoque`.
- Map port `5432` of the container to port `5432` on your host machine.

## Accessing the Database

You can connect to the PostgreSQL database using any PostgreSQL client with the following credentials:

- Host: `localhost`
- Port: `5432`
- User: `postgres`
- Password: `123456`
- Database: `estoque`

## Stopping and Removing the Container

To stop the running container, use:

```
docker stop postgres-go
```

To remove the container, use:

```
docker rm postgres-go
```

## License

This project is licensed under the MIT License.