package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"tiktok/conf"
	"tiktok/dao"
)

func main() {
	fmt.Print(
		"                                                 \n" +
			"███╗   ██╗ ██████╗     ██████╗ ██╗   ██╗ ██████╗ \n" +
			"████╗  ██║██╔═══██╗    ██╔══██╗██║   ██║██╔════╝ \n" +
			"██╔██╗ ██║██║   ██║    ██████╔╝██║   ██║██║  ███╗\n" +
			"██║╚██╗██║██║   ██║    ██╔══██╗██║   ██║██║   ██║\n" +
			"██║ ╚████║╚██████╔╝    ██████╔╝╚██████╔╝╚██████╔╝\n" +
			"╚═╝  ╚═══╝ ╚═════╝     ╚═════╝  ╚═════╝  ╚═════╝ \n" +
			"                                                 \n")
	//初始化setting
	conf.Init()
	//使用gin的默认http配置
	r := gin.Default()

	//初始化路由
	initRouter(r)
	dao.InitMysql()
	dao.InitTables()

	err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		panic("服务器启动失败")
	}

}
