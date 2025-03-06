package repo

import (
	"context"
	kztv1 "github.com/JohnBilbo/transfer_kzt_proto/gen/go/transfer_kzt_grpc"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repo interface {
}

type repo struct {
	pgxPool *pgxpool.Pool
}

func NewRepo(pgxPool *pgxpool.Pool) Repo {
	return &repo{
		pgxPool: pgxPool,
	}
}

func (r *repo) TransferByID(ctx context.Context, id int64) (*kztv1.TransferKZT, error) {
	//r.pgxPool.Query(ctx, "SELECT COUNT(*) FROM transfers WHERE id = $1", id)
	return nil, nil
}
