package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"user_service_sso/internal/app"
	"user_service_sso/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// инициализация конфига
	cfg := config.MustLoad()

	// инициализация логгера
	log := setupLogger(cfg.Env)

	log.Info("starting application")

	// инициализация приложения (app)
	application := app.New(log, cfg.GRPC.Port, cfg.DB, cfg.TokenTTL)

	// запустить gRPC-сервер приложения
	go application.GRPCSrv.MustRun()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	signal := <-stop

	log.Info("stopping application", slog.String("signal", signal.String()))

	application.GRPCSrv.Stop()
	log.Info("application stopped")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
