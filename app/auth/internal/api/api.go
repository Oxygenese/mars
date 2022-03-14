package api

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/mars-projects/oauth2/v4/server"
)

var ProviderApiSet = wire.NewSet(NewTokenHandler, NewCaptchaHandler)

type TokenApi struct {
	server *server.Server
}

func NewTokenHandler(server *server.Server) *TokenApi {
	return &TokenApi{server: server}
}

type CaptchaApi struct {
	log *log.Helper
}

func NewCaptchaHandler(logger log.Logger) *CaptchaApi {
	return &CaptchaApi{log: log.NewHelper(logger)}
}
