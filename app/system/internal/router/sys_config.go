package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mars-projects/mars/app/system/internal/api"
)

func init() {
	routerPublicRole = append(routerPublicRole, registerSysConfigPublicRouter)
	routerAuthentication = append(routerAuthentication, registerSysConfigRouter)
}

func registerSysConfigRouter(r *gin.RouterGroup, option *api.ApiOption) {
	h := option.SysConfigHandler
	g := r.Group("config")
	{
		g.GET("/:id", h.Get)
		g.PUT("/:id", h.Update)
		g.POST("", h.Insert)
		g.DELETE("", h.Delete)
		g.GET("/page", h.GetPage)
		g.GET("", h.Get2Set)
		g.PUT("", h.Update2Set)
		g.GET("/configKey/:configKey", h.GetSysConfigByKEYForService)
	}
}
func registerSysConfigPublicRouter(r *gin.RouterGroup, option *api.ApiOption) {
	h := option.SysConfigHandler
	g := r.Group("/app-config")
	{
		g.GET("", h.Get2SysApp)
	}
}
