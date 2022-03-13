package client

import (
	"context"
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	"github.com/mars-projects/mars/api/master"
	"github.com/mars-projects/mars/conf"
	"log"
)

var ProviderMasterClientSet = wire.NewSet(NewMasterClient)

func NewMasterClient(r *nacos.Registry, client *conf.Client) (master.MasterClient, error) {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(grpcClientAddr(client.Master)),
		grpc.WithDiscovery(r),
	)
	if err != nil {
		log.Println(err)
	}
	return master.NewMasterClient(conn), err
}
