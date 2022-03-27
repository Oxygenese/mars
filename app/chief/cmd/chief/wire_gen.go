// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/mars-projects/mars/app/chief/internal/server"
	"github.com/mars-projects/mars/app/chief/internal/service"
	"github.com/mars-projects/mars/app/chief/internal/task"
	"github.com/mars-projects/mars/common/wire/data"
	"github.com/mars-projects/mars/common/wire/register"
	"github.com/mars-projects/mars/common/wire/sender"
	"github.com/mars-projects/mars/common/wire/transaction"
	"github.com/mars-projects/mars/conf"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, logger log.Logger, registry *conf.Registry) (*kratos.App, func(), error) {
	client, cleanup, err := data.NewRedisClient(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	tokenStore := data.NewRedisTokenStore(client)
	transactionSender := sender.NewMessageSender()
	engine := transaction.NewTransactionEngine(transactionSender)
	tasksManager := task.NewTaskManager(engine)
	chiefService := service.NewChiefService(tasksManager)
	httpServer := server.NewHTTPServer(confServer, tokenStore, chiefService, logger)
	grpcServer := server.NewGRPCServer(confServer, chiefService, logger)
	nacosRegistry := register.NewNacosRegistrar(registry)
	app := newApp(logger, httpServer, grpcServer, confServer, nacosRegistry)
	return app, func() {
		cleanup()
	}, nil
}
