package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"transfer_kzt_grpc/internal/app"
	"transfer_kzt_grpc/internal/config"
	"transfer_kzt_grpc/internal/lib/logger"
	"transfer_kzt_grpc/internal/storage/postgre"
)

func main() {
	const op = "main"
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cfg := config.NewConfig()
	log := logger.NewLogger(cfg)
	defer log.Sync()

	pool, err := postgre.NewPool(ctx, cfg)
	if err != nil {
		log.Sugar().With("op", op).Error(err)
		pool.Close()
		panic(err)
	}
	defer pool.Close()

	// svc
	application := app.NewApp(log.Sugar(), cfg)
	go application.GRPCSvc.MustRun()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGTERM)
	stopInfo := <-stop
	log.Sugar().With("op", op, "port", cfg.GRPCConfig.Port).Info(fmt.Sprintf("server stop info: %s", stopInfo.String()))
	application.GRPCSvc.Stop(pool, log, cancel)
	log.Sugar().With("op", op, "port", cfg.GRPCConfig.Port).Info("grpc server successfully stopped")
}
