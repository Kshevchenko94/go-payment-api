package main

import (
	"log"
	"log/slog"
	"os"

	"Kshevchenko94/go-payment-api/internal/infrastructure/storage"

	"github.com/joho/godotenv"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	dbCfg := storage.PostgresConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", ""),
		DBName:   getEnv("DB_NAME", "postgres"),
		SSLMode:  "disable",
	}

	db, err := storage.DialPostgres(dbCfg)
	if err != nil {
		logger.Error("failed to connect to postgres", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	logger.Info("successfully connected to postgres pool")

	// 3. Сюда мы дальше передадим переменную `db` в репозитории:
	// userRepo := storage.NewUserRepository(db)
	// ...

	logger.Info("starting http server on :8080")
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
