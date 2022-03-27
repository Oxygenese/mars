package main

import (
	"flag"
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/mars-projects/mars/conf"
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, hs *http.Server, gs *grpc.Server, server *conf.Server, rr *nacos.Registry) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(server.Name),
		kratos.Version(server.Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Registrar(rr),
		kratos.Server(
			hs,
			gs,
		),
	)
}

func main() {
	flag.Parse()

	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	var rc conf.Registry
	if err := c.Scan(&rc); err != nil {
		panic(err)
	}
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", bc.Server.Name,
		"service.version", bc.Server.Version,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
	)
	app, cleanup, err := initApp(bc.Server, bc.Data, logger, &rc)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
