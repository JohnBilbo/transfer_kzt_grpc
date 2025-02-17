package app

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// App - структура приложения
type App struct {
	log        *zap.SugaredLogger
	gRPCServer *grpc.Server
	port       int
}

func New(log *zap.SugaredLogger) *App {
	gRPCServer := grpc.NewServer()
	return &App{}
}
