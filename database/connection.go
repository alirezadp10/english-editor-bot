package database

import (
    "fmt"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "os"
)

func ConnectDB() (*gorm.DB, error) {
    dsn := fmt.Sprintf(
        "host=postgres user=%s password=%s dbname=%s port=%s sslmode=disable",
        os.Getenv("POSTGRES_USER"),
        os.Getenv("POSTGRES_PASSWORD"),
        os.Getenv("POSTGRES_DB"),
        os.Getenv("POSTGRES_PORT"),
    )
    return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
