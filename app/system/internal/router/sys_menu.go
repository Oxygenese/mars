package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mars-projects/mars/app/system/internal/api"
)

func init() {
	routerAuthentication = append(routerAuthentication, registerSysMenuRouter)
}

func registerSysMenuRouter(r *gin.RouterGroup, option *api.ApiOption) {
	h := option.SysMenuHandler
	g := r.Group("/menu")
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
