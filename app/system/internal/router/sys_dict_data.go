package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mars-projects/mars/app/system/internal/api"
)

func init() {
	routerAuthentication = append(routerAuthentication, registerSysDictDataRouter)
}

func registerSysDictDataRouter(r *gin.RouterGroup, option *api.ApiOption) {
	h := option.SysDictDataHandler
	g := r.Group("/dict/data")
	{
		g.GET("/page", h.GetPage)
		g.GET("/:dictCode", h.Get)
		g.POST("", h.Insert)
		g.PUT("/:dictCode", h.Update)
		g.DELETE("", h.Delete)
		g.GET("/option-select", h.GetAll)
	}
}
