package app

import (
	"log/slog"
	"time"
	grpcapp "user_service_sso/internal/app/grpc"
	"user_service_sso/internal/config"
	"user_service_sso/internal/services/auth"
	"user_service_sso/internal/storage/postgres"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	dbConfig config.DBConfig,
	tokenTTL time.Duration,
) *App {
	// инициализировать хранилище
	storage, err := postgres.New(dbConfig)
	if err != nil {
		panic(err)
	}

	// инициализировать grpc service layer
	authService := auth.New(log, storage, tokenTTL)

	grpcApp := grpcapp.New(log, *authService, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
