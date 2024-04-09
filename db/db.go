package db

import (
    "database/sql"
    "fmt"
    "os"

    _ "github.com/lib/pq" // Import PostgreSQL driver

    "github.com/joho/godotenv"
)

// Connect to the database
func Connect() (*sql.DB, error) {
    if err := godotenv.Load(); err != nil {
        return nil, fmt.Errorf("failed to load environment variables: %w", err)
    }

    connStr := os.Getenv("DATABASE_URL")
    if connStr == "" {
        return nil, fmt.Errorf("missing DATABASE_URL environment variable")
    }

    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, fmt.Errorf("failed to connect to database: %w", err)
    }

    if err := db.Ping(); err != nil {
        return nil, fmt.Errorf("failed to ping database: %w", err)
    }

    // Successful connection established
    fmt.Println("Connected to database successfully!")

    return db, nil
}

// Get the database version
func GetDBVersion(db *sql.DB) (string, error) {
    var version string
    err := db.QueryRow("select version()").Scan(&version)
    if err != nil {
        return "", fmt.Errorf("failed to query database version: %w", err)
    }
    return version, nil
}
