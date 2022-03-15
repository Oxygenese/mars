package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/mars-projects/mars/app/auth/internal/api"
	"github.com/mars-projects/mars/conf"
	https "net/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, logger log.Logger, captcha *api.CaptchaApi, token *api.TokenApi) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(),
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
	r := srv.Route("/auth", CorsFilter)
	r.POST("/token", token.Token)
	r.GET("/confirm", token.Confirm)
	r.GET("/refresh_token", token.RefreshToken)
	r.GET("/authorize", token.Authorize)
	r.GET("captcha", captcha.GenerateCaptchaHandler)
	r.PUT("logout", token.Logout)
	return srv
}
func CorsFilter(next https.Handler) https.Handler {
	return https.HandlerFunc(func(w https.ResponseWriter, req *https.Request) {
		ctx := req.Context()
		info, _ := transport.FromServerContext(ctx)
		info.ReplyHeader().Set("Access-Control-Allow-Origin", "*")
		info.ReplyHeader().Set("Access-Control-Allow-Credentials", "true")
		info.ReplyHeader().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
		info.ReplyHeader().Set("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session")
		info.ReplyHeader().Set("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")

		next.ServeHTTP(w, req.WithContext(ctx))
	})
}
