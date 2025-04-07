package main

import (
	"database/sql"
	"fmt"
	"log"

	"os"

	"github.com/joho/godotenv"
	"github.com/ramonxm/gatewayapp/gateway-api/internal/repository"
	"github.com/ramonxm/gatewayapp/gateway-api/internal/service"
	"github.com/ramonxm/gatewayapp/gateway-api/internal/web/server"

	_ "github.com/lib/pq"
)

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)

	if len(value) == 0 {
		return defaultValue
	}

	return value
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_NAME", "postgres"),
		getEnv("DB_SSLMODE", "disable"))

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	accountRepository := repository.NewAccountRepository(db)
	accountService := service.NewAccountService(accountRepository)

	port := getEnv("PORT", "8080")

	srv := server.NewServer(port, accountService)
	srv.ConfigureRoutes()

	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
