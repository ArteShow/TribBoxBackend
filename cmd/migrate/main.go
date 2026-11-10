package main

import (
	"context"
	"log"
	"os"

	"github.com/ArteShow/TribBoxBackend/internal/database"
	"github.com/ArteShow/TribBoxBackend/internal/migrate"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
	log.Println("Successfully loaded .env")

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("connString is empty")
	}
	log.Println("Successfully loaded the connString from .env")

	db, err := database.Connect(context.Background(), dbURL)
	if err != nil {
		panic(err)
	}
	defer db.Conn.Close()

	log.Println("Successfully connected to the db")

	if err = migrate.RunMigrations(context.Background(), &db); err != nil {
		panic(err)
	}
	log.Println("Successfully executed migrations")
}
