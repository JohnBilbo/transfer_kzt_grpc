package transfer_kzt

import (
	"context"
	kztv1 "github.com/JohnBilbo/transfer_kzt_proto/gen/go/transfer_kzt_grpc"
	"google.golang.org/grpc"
)

type transferKZTServerAPI struct {
	kztv1.UnimplementedTransferKZTServiceServer
}

// Register - регистрирует функции
func Register(gRPC *grpc.Server) {
	kztv1.RegisterTransferKZTServiceServer(gRPC, &transferKZTServerAPI{})
}

func (t *transferKZTServerAPI) Transfer(ctx context.Context, kzt *kztv1.RequestTransfer) (*kztv1.ResponseTransfer, error) {
	a := &kztv1.ResponseTransfer{
		Transfer: &kztv1.TransferKZT{
			Id: 10,
		},
	}
	return a, nil
}

func (t *transferKZTServerAPI) CreateTransfer(ctx context.Context, kzt *kztv1.RequestCreateTransfer) (*kztv1.ResponseCreateTransfer, error) {
	a := &kztv1.ResponseCreateTransfer{
		Status: "Resp",
	}
	return a, nil
}
