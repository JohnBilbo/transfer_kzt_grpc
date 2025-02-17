package transfer_kzt_server

import (
	"context"
	t_kzt_v1 "github.com/JohnBilbo/transfer_kzt_proto/gen/go/transfer_kzt_grpc"
	"google.golang.org/grpc"
)

type transferKZTServerAPI struct {
	t_kzt_v1.UnimplementedTransferKZTServiceServer
}

// Register - регистрирует сервс
func Register(gRPC *grpc.Server) {
	t_kzt_v1.RegisterTransferKZTServiceServer(gRPC, &transferKZTServerAPI{})
}

func (t *transferKZTServerAPI) Transfer(ctx context.Context, kzt *t_kzt_v1.RequestTransfer) (*t_kzt_v1.ResponseTransfer, error) {
	return nil, nil
}

func (t *transferKZTServerAPI) CreateTransfer(ctx context.Context, kzt *t_kzt_v1.RequestTransfer) (*t_kzt_v1.ResponseTransfer, error) {
	return nil, nil
}
