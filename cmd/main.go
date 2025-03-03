package main

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
	"os/signal"
	"syscall"
	"transfer_kzt_grpc/internal/app"
	"transfer_kzt_grpc/internal/config"
	"transfer_kzt_grpc/internal/lib/logger"
)

func main() {
	const op = "main"
	cfg := config.DefaultConfig()
	log := logger.NewLogger(cfg)
	defer log.Sync()

	// db
	pool := &pgxpool.Pool{}

	application := app.NewApp(log.Sugar(), cfg)
	go application.GRPCSvc.MustRun()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGTERM)
	stopInfo := <-stop
	log.Sugar().With("op", op, "port", cfg.GRPCConfig.Port).Info(fmt.Sprintf("server stop info: %s", stopInfo.String()))
	application.GRPCSvc.Stop(pool)
	log.Sugar().With("op", op, "port", cfg.GRPCConfig.Port).Info("grpc server successfully stopped")
}
