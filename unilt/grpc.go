package unilt

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type Grpc struct {
	Address string
}

func NewGrpc(address string) Grpc {
	return Grpc{Address: address}
}

func (g *Grpc) RegisterGrpcApi() *grpc.ClientConn {
	conn, err := grpc.NewClient(g.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("没有连接: %v", err)
	}
	return conn
}
