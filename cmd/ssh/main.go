package main

import (
	"github.com/iAbbos/go-grpc_auth-sso/internal/app"
	"github.com/iAbbos/go-grpc_auth-sso/internal/config"
	"github.com/iAbbos/go-grpc_auth-sso/internal/pkg/lib/logger/handlers/slogpretty"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	//Implemented config
	cfg := config.MustLoad()

	//Implemented logger
	log := setupLogger(cfg.Env)

	log.Info("Starting application", slog.Any("cfg", cfg))

	//Implemented application
	application := app.New(log, cfg.GRPConfig.Port, cfg.StoragePath, cfg.TokenTTL)
	go application.GRPCSrv.MustRun()

	//TODO: Implement grpc server

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop

	log.Info("Shutting down application", slog.Any("signal", sign.String()))

	application.GRPCSrv.Stop()

	log.Info("Application stopped")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = setupPrettySlog()
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

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{Level: slog.LevelDebug},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
