package main

import (
	"blog_server/core"
	"blog_server/global"
	"fmt"
)

func main() {
	core.InitConf() //配置文件初始化
	global.DB = core.InitGorm()
	fmt.Print(global.DB)
}
