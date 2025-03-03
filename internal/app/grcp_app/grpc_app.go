package grcp_app

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"transfer_kzt_grpc/internal/config"
	"transfer_kzt_grpc/internal/grpc/transfer_kzt"
	"transfer_kzt_grpc/internal/lib/c_errors"
)

// App - структура приложения
type App struct {
	log     *zap.SugaredLogger
	gRPCSvc *grpc.Server
	port    int
}

func NewApp(log *zap.SugaredLogger, cfg *config.Config) *App {
	gRPCSvc := grpc.NewServer()
	registerGRPC(gRPCSvc)
	return &App{
		log:     log,
		gRPCSvc: gRPCSvc,
		port:    cfg.GRPCConfig.Port,
	}
}

// registerGRPC - здесь регистрируем все наши grpc функции
func registerGRPC(gRPCServer *grpc.Server) {
	transfer_kzt.Register(gRPCServer)
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "gRPC.Run"
	a.log.With("op", op, "port", a.port).Info("starting app")
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		a.log.With("op", op).Error(err)
		return c_errors.AttemptToRunErr
	}
	a.log.With("op", op, "port", a.port).Info("starting grpc server")
	if err = a.gRPCSvc.Serve(l); err != nil {
		a.log.With("op", op).Error(err)
		return c_errors.AttemptToRunErr
	}
	return nil
}

func (a *App) Stop(pool *pgxpool.Pool) {
	const op = "gRPC.Stop"
	a.log.With("op", op, "port", a.port).Info("stop grpc server")
	a.gRPCSvc.GracefulStop()
	a.log.With("op", op, "port", a.port).Info("close postgres pool")
	pool.Close()
}
