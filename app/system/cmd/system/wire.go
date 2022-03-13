//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/mars-projects/mars/app/system/internal/api"
	"github.com/mars-projects/mars/app/system/internal/biz"
	"github.com/mars-projects/mars/app/system/internal/router"
	"github.com/mars-projects/mars/app/system/internal/server"
	"github.com/mars-projects/mars/app/system/internal/service"
	"github.com/mars-projects/mars/conf"
	"github.com/mars-projects/mars/lib/wire/data"
	"github.com/mars-projects/mars/lib/wire/middleware/oauth"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		api.ProviderApiOptionSet,
		data.ProviderRedisTokenStoreSet,
		data.ProviderRedisSet,
		router.ProviderEngineSet,
		oauth.ProviderAuthenticationSet,
		data.ProviderGormSet,
		server.ProviderSet,
		biz.ProviderBizSet,
		service.ProviderSet,
		newApp))
}
