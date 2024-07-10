package routers

import (
	"blog_server/middleware"
	"blog_server/models/res"

	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	router := gin.Default()
	res.InitErrorCode()
	router.Static("/res", "./res")
	SignApiRouterGroup := router.Group("api")
	NoSignApiRouterGroup := router.Group("api")
	noSignRouterGroupApp := RouterGroup{NoSignApiRouterGroup}
	routerGroupApp := RouterGroup{SignApiRouterGroup}
	routerGroupApp.Use(middleware.JWTAuth())
	routerGroupApp.SettingRouter()
	routerGroupApp.MarkdownhtmRouter()
	routerGroupApp.DocumentRouter()
	noSignRouterGroupApp.UserRouter()
	return router
}
