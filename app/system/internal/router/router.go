package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/mars-projects/mars/app/system/internal/api"
	"github.com/mars-projects/mars/lib/wire/middleware/cors"
	"github.com/mars-projects/mars/lib/wire/middleware/oauth"
)

var ProviderEngineSet = wire.NewSet(NewGinEngine)

func NewGinEngine(authentication *oauth.Authentication, option *api.ApiOption) *gin.Engine {
	r := gin.Default()
	r.Use(cors.CorsFunc())
	InitSysRouter(r, option, authentication)
	return r
}

var (
	routerPublicRole     = make([]func(*gin.RouterGroup, *api.ApiOption), 0)
	routerAuthentication = make([]func(r *gin.RouterGroup, h *api.ApiOption, authMiddleware *oauth.Authentication), 0)
)

func InitSysRouter(r *gin.Engine, option *api.ApiOption, authMiddleware *oauth.Authentication) {
	// 无需认证的路由
	publicRouter(r, option)
	// 需要认证的路由
	authenticationRouter(r, option, authMiddleware)
}

// 无需认证的路由示例
func publicRouter(r *gin.Engine, option *api.ApiOption) {
	// 可根据业务需求来设置接口版本
	v1 := r.Group("sys")
	for _, f := range routerPublicRole {
		f(v1, option)
	}
}

// 需要认证的路由示例
func authenticationRouter(r *gin.Engine, option *api.ApiOption, authMiddleware *oauth.Authentication) {
	// 可根据业务需求来设置接口版本
	v1 := r.Group("sys")
	for _, f := range routerAuthentication {
		f(v1, option, authMiddleware)
	}
}
