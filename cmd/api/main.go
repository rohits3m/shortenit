package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"shortenit/internal/models"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

type Config struct {
	env     string
	port    string
	dbStr   string
	baseUrl string
}

type Application struct {
	db     *pgx.Conn
	logger *slog.Logger
	config *Config

	// Models
	links *models.LinkModel
}

func main() {
	// Configuration
	godotenv.Load()
	config := Config{
		env:     os.Getenv("ENV"),
		port:    os.Getenv("PORT"),
		dbStr:   os.Getenv("DB_URL"),
		baseUrl: os.Getenv("BASE_URL"),
	}

	// Logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))

	// Database Connection with Timeout
	conn, err := pgx.Connect(context.Background(), config.dbStr)
	if err != nil {
		logger.Error(err.Error())
	}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = conn.Ping(ctxTimeout)
	if err != nil {
		logger.Error(err.Error())
	}

	// Application Instance
	app := Application{
		db:     conn,
		logger: logger,
		config: &config,
		links:  &models.LinkModel{Db: conn},
	}

	// Server
	server := http.Server{
		Addr:         fmt.Sprintf(":%s", config.port),
		Handler:      app.routes(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("Server starting", "port", config.port, "env", config.env)
	err = server.ListenAndServe()
	if err != nil {
		logger.Error(err.Error())
	}
}
