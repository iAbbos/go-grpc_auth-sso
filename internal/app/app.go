package app

import (
	grpcapp "github.com/iAbbos/go-grpc_auth-sso/internal/app/grpc"
	"github.com/iAbbos/go-grpc_auth-sso/internal/domain/repository/sqlite"
	"github.com/iAbbos/go-grpc_auth-sso/internal/services/auth"
	"log/slog"
	"time"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {
	repo, err := sqlite.New(storagePath)
	if err != nil {
		panic(err)
	}

	authService := auth.New(log, repo, repo, repo, tokenTTL)

	grpcApp := grpcapp.New(log, authService, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}

}
