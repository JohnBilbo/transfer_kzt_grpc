package app

import (
	"go.uber.org/zap"
	"transfer_kzt_grpc/internal/app/grcp_app"
	"transfer_kzt_grpc/internal/config"
)

type App struct {
	GRPCSvc *grcp_app.App
}

func NewApp(log *zap.SugaredLogger, cfg *config.Config) *App {
	grpcApp := grcp_app.NewApp(log, cfg)
	return &App{
		GRPCSvc: grpcApp,
	}
}
