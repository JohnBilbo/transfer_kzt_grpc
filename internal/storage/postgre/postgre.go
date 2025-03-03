package postgre

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"transfer_kzt_grpc/internal/config"
)

func NewPool(ctx context.Context, cfg *config.Config) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		cfg.DBConfig.Host,
		cfg.DBConfig.User,
		cfg.DBConfig.Password,
		cfg.DBConfig.Name,
		cfg.DBConfig.Port,
		cfg.DBConfig.SSLMode,
		cfg.DBConfig.TimeZone,
	)
	dbCfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}
	dbCfg.MaxConns = cfg.DBConfig.MaxConns
	dbCfg.MaxConnLifetime = cfg.DBConfig.MaxConnLifeTime
	dbCfg.MaxConnLifetimeJitter = cfg.DBConfig.MaxConnLifeTimeJitter
	dbCfg.MaxConnIdleTime = cfg.DBConfig.MaxConnIdleTime
	pool, err := pgxpool.NewWithConfig(ctx, dbCfg)
	if err != nil {
		return nil, err
	}
	return pool, nil
}
