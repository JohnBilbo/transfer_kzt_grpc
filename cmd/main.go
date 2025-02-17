package main

import (
	grpcS "github.com/JohnBilbo/transfer_kzt_proto/gen/go/transfer_kzt_grpc"
	"os"
	"transfer_kzt_grpc/internal/config"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	c, err := config.FindProjectRoot(cwd)
	if err != nil {
		panic(err)
	}
	print(c)
	grpcS.UnimplementedTransferKZTServiceServer{}
}
