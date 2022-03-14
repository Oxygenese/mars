package client

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/mars-projects/mars/api/system"
	"github.com/mars-projects/mars/conf"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	"log"
)

var ProviderSysClientSet = wire.NewSet(NewSysClient)

func NewSysClient(r *nacos.Registry, client *conf.Client) (system.SystemClient, error) {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(grpcClientAddr(client.Sys)),
		grpc.WithDiscovery(r),
	)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	return system.NewSystemClient(conn), nil
}

func grpcClientAddr(serveName string) string {
	return fmt.Sprintf("discovery:///%s.grpc", serveName)
}
