package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mars-projects/mars/app/system/internal/api"
)

func init() {
	routerAuthentication = append(routerAuthentication, registerSysDictTypeRouter)
}

func registerSysDictTypeRouter(r *gin.RouterGroup, option *api.ApiOption) {
	h := option.SysDictTypeHandler
	g := r.Group("/dict/type")
	{
		g.GET("/page", h.GetPage)
		g.GET("/:id", h.Get)
		g.POST("", h.Insert)
		g.PUT("/:id", h.Update)
		g.DELETE("", h.Delete)
		g.GET("/option-select", h.GetAll)
	}
}
