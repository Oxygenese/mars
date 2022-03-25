package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mars-projects/mars/app/system/internal/api"
)

func init() {
	routerAuthentication = append(routerAuthentication, registerSysDeptRouter)
}

func registerSysDeptRouter(r *gin.RouterGroup, option *api.ApiOption) {
	h := option.SysDeptHandler
	g := r.Group("/dept")
	{
		g.GET("/page", h.GetPage)
		g.GET("/:id", h.Get)
		g.POST("", h.Insert)
		g.PUT("/:deptId", h.Update)
		g.DELETE("", h.Delete)
		g.GET("/tree", h.Get2Tree)
		g.GET("/tree/dept/select/:roleId", h.GetDeptTreeRoleSelect)
	}
}
