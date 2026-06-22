package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect initializes the database connection.
// It supports two modes based on the environment variable DB_TYPE:
//   - "postgres": connects to a PostgreSQL instance (default behavior).
//   - "sqlite": uses a local SQLite file (plug‑and‑play for Windows users).
func Connect() {
	// Load .env file if present.
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file loaded, relying on environment variables")
	}

	dbType := os.Getenv("DB_TYPE")
	if dbType == "sqlite" {
		// Use SQLite file stored in the project root.
		sqlitePath := os.Getenv("SQLITE_PATH")
		if sqlitePath == "" {
			sqlitePath = "stock_manager.db"
		}
		var err error
        DB, err = gorm.Open(sqlite.Open(sqlitePath), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to open SQLite database: %v", err)
		}
		fmt.Println("SQLite database connected at", sqlitePath)
		return
	}

	// Default to PostgreSQL.
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar no banco PostgreSQL: %v", err)
	}

	DB = database

	fmt.Println("PostgreSQL conectado")
}
