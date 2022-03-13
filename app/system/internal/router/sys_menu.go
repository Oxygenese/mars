package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mars-projects/mars/app/system/internal/api"
	"github.com/mars-projects/mars/lib/wire/middleware/oauth"
)

func init() {
	routerAuthentication = append(routerAuthentication, registerSysMenuRouter)
}

func registerSysMenuRouter(r *gin.RouterGroup, option *api.ApiOption, authMiddleware *oauth.Authentication) {
	h := option.SysMenuHandler
	g := r.Group("/menu").Use(authMiddleware.GinAuthMiddleware())
	{
		g.GET("/menurole", h.GetMenuRole)
		g.GET("/page", h.GetPage)
		g.GET("/:id", h.Get)
		g.GET("/tree/select/:id", h.GetMenuTreeSelect)
		g.POST("", h.Insert)
		g.PUT("/:id", h.Update)
		g.DELETE("", h.Delete)
	}
}
