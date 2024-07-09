package routers

import (
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
	apiRouterGroup := router.Group("api")
	routerGroupApp := RouterGroup{apiRouterGroup}
	routerGroupApp.SettingRouter()
	routerGroupApp.MarkdownhtmRouter()
	routerGroupApp.DocumentRouter()
	routerGroupApp.UserRouter()
	return router
}
