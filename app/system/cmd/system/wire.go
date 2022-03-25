//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/mars-projects/mars/app/system/internal/biz"
	"github.com/mars-projects/mars/app/system/internal/server"
	"github.com/mars-projects/mars/app/system/internal/service"
	"github.com/mars-projects/mars/app/system/internal/task"
	"github.com/mars-projects/mars/common/wire/data"
	"github.com/mars-projects/mars/common/wire/register"
	"github.com/mars-projects/mars/common/wire/sender"
	"github.com/mars-projects/mars/common/wire/transaction"
	"github.com/mars-projects/mars/conf"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Registry, *conf.Data, *conf.Auth, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ProviderServerSet,
		service.ProviderServiceSet,
		register.ProviderNacosSet,
		task.ProviderTasksManagerSet,
		transaction.ProviderTransactionEngineSet,
		sender.ProviderMessageSenderSet,
		data.ProviderRedisTokenStoreSet,
		data.ProviderRedisSet,
		biz.ProviderBizSet,
		data.ProviderGormSet,
		newApp))
}
