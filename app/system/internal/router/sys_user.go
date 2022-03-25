package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mars-projects/mars/app/system/internal/api"
)

func init() {
	routerAuthentication = append(routerAuthentication, registerSysUserRouter)
}
func registerSysUserRouter(r *gin.RouterGroup, option *api.ApiOption) {
	h := option.SysUserHandler
	g := r.Group("/user")
	{
		g.GET("/:id", h.Get)
		g.GET("/page", h.GetPage)
		g.POST("/create", h.Insert)
		g.PUT("", h.Update)
		g.DELETE("/delete", h.Delete)
		g.PUT("/status", h.UpdateStatus)
		g.PUT("/pwd/reset", h.ResetPwd)
		g.POST("/avatar", h.InsetAvatar)
		g.GET("/profile", h.GetProfile)
		g.PUT("/pwd/set", h.UpdatePwd)
		g.GET("/info", h.GetInfo)
	}
}
