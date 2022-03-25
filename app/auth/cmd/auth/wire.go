//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/mars-projects/mars/app/auth/internal/api"
	"github.com/mars-projects/mars/app/auth/internal/biz"
	"github.com/mars-projects/mars/app/auth/internal/oauth2"
	"github.com/mars-projects/mars/app/auth/internal/server"
	"github.com/mars-projects/mars/common/wire/client"
	"github.com/mars-projects/mars/common/wire/data"
	"github.com/mars-projects/mars/common/wire/register"
	"github.com/mars-projects/mars/conf"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Registry, *conf.Data, *conf.Auth, log.Logger, *conf.Client) (*kratos.App, func(), error) {
	panic(wire.Build(
		api.ProviderApiSet,
		oauth2.ProviderOauth,
		data.ProviderRedisSet,
		data.ProviderRedisTokenStoreSet,
		server.ProviderSet,
		client.ProviderSysClientSet,
		biz.ProviderSet,
		register.ProviderNacosSet,
		newApp))
}
