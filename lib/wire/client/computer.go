package client

import (
	"context"
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	"github.com/mars-projects/mars/api/compute"
	"github.com/mars-projects/mars/conf"
	"log"
)

var ProviderComputerClientSet = wire.NewSet(NewComputerClient)

func NewComputerClient(r *nacos.Registry, client *conf.Client) (compute.ComputeClient, error) {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(grpcClientAddr(client.Computer)),
		grpc.WithDiscovery(r),
	)
	if err != nil {
		log.Println(err)
	}
	return compute.NewComputeClient(conn), err
}
