package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/khudob1n/krasnodar-parking/backend/internal/httpapi"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	addr := ":" + envOr("PORT", "8080")

	srv := &http.Server{
		Addr:              addr,
		Handler:           httpapi.NewRouter(),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	logger.Info("Сервер запущен", "addr", addr)
	if err := srv.ListenAndServe(); err != nil {
		logger.Error("Сервер остановился", "err", err)
		os.Exit(1)
	}
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
