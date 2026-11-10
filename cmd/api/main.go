package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"Kshevchenko94/go-payment-api/internal/infrastructure/storage"
	"Kshevchenko94/go-payment-api/internal/infrastructure/transport/http/delivery"
	"Kshevchenko94/go-payment-api/internal/usecase"

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
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
	}

	db, err := storage.DialPostgres(dbCfg)
	if err != nil {
		logger.Error("failed to connect to postgres pool", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	logger.Info("successfully connected to postgres pool")

	// В будущем переменная db передастся в репозитории:
	userRepo := storage.NewUserRepository(db)
	authUseCase := usecase.NewAuthInteractor(userRepo)
	authHandler := delivery.NewAuthHandler(authUseCase)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/v1/auth/register", authHandler.Register)

	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}

	logger.Info("starting http server on :8080")
	err = s.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logger.Error("server failed", "error", err)
		os.Exit(1)
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
