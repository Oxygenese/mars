package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/mars-projects/mars/api"
	"github.com/mars-projects/mars/app/chief/internal/service"
	"github.com/mars-projects/mars/common/middleware/authentication"
	"github.com/mars-projects/mars/conf"
	"github.com/mars-projects/oauth2/v4"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, store oauth2.TokenStore, service *service.ChiefService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			logging.Server(logger),
			authentication.Server(store, logger),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	api.RegisterChiefHTTPServer(srv, service)
	return srv
}
