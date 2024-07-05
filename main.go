package main

import (
	"blog_server/core"
	"blog_server/global"
	"blog_server/routers"
	"fmt"
	"time"
)

func main() {
	core.InitConf()                //配置文件初始化
	global.LOG = core.InitLogger() //初始化日志
	global.DB = core.InitGorm()    //连接数据库
	router := routers.InitRouter() //初始化路由
	fmt.Println(time.Now())
	router.Run("127.0.0.1:8081") //路由端口

}
