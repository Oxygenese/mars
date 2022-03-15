package register

import (
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/google/wire"
	"github.com/mars-projects/mars/conf"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"log"
)

// ProviderNacosSet is compute providers.
var ProviderNacosSet = wire.NewSet(NewNacosRegistrar)

func NewNacosRegistrar(registry *conf.Registry) *nacos.Registry {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(registry.Nacos.Address, registry.Nacos.Port),
	}
	cc := constant.NewClientConfig(
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("./nacos/log"),
		constant.WithCacheDir("./cache"),
		constant.WithLogLevel("error"),
	)
	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ServerConfigs: sc,
			ClientConfig:  cc,
		},
	)
	if err != nil {
		log.Panic(err)
	}
	return nacos.New(client)
}
