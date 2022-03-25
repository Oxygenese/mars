package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/mars-projects/mars/app/system/internal/api"
)

var ProviderEngineSet = wire.NewSet(NewGinEngine)

func NewGinEngine(option *api.ApiOption) *gin.Engine {
	r := gin.Default()
	InitSysRouter(r, option)
	return r
}

var (
	routerPublicRole     = make([]func(*gin.RouterGroup, *api.ApiOption), 0)
	routerAuthentication = make([]func(r *gin.RouterGroup, h *api.ApiOption), 0)
)

func InitSysRouter(r *gin.Engine, option *api.ApiOption) {
	// 无需认证的路由
	publicRouter(r, option)
	// 需要认证的路由
	authenticationRouter(r, option)
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
func authenticationRouter(r *gin.Engine, option *api.ApiOption) {
	// 可根据业务需求来设置接口版本
	v1 := r.Group("sys")
	for _, f := range routerAuthentication {
		f(v1, option)
	}
}
