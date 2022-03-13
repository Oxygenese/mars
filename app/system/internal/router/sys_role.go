package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mars-projects/mars/app/system/internal/api"
	"github.com/mars-projects/mars/lib/wire/middleware/oauth"
)

func init() {
	routerAuthentication = append(routerAuthentication, registerSysRoleRouter)
}

func registerSysRoleRouter(r *gin.RouterGroup, option *api.ApiOption, authMiddleware *oauth.Authentication) {
	h := option.SysRoleHandler
	g := r.Group("role").Use(authMiddleware.GinAuthMiddleware())
	{
		g.GET("/page", h.GetPage)
		g.GET("/:id", h.Get)
		g.POST("", h.Insert)
		g.PUT("/:id", h.Update)
		g.PUT("/status/:id", h.Update2Status)
		g.PUT("/data/scope", h.Update2DataScope)
		g.DELETE("", h.Delete)
	}
}
