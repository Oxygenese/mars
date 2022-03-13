package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mars-projects/mars/app/system/internal/api"
	"github.com/mars-projects/mars/lib/wire/middleware/oauth"
)

func init() {
	routerAuthentication = append(routerAuthentication, registerSysPostRouter)
}

func registerSysPostRouter(r *gin.RouterGroup, option *api.ApiOption, authMiddleware *oauth.Authentication) {
	h := option.SysPostHandler
	g := r.Group("post").Use(authMiddleware.GinAuthMiddleware())
	{
		g.GET("/page", h.GetPage)
		g.GET("/:id", h.Get)
		g.POST("", h.Insert)
		g.PUT("/:id", h.Update)
		g.DELETE("", h.Delete)
	}
}
