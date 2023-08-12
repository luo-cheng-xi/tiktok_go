package main

import (
	"github.com/gin-gonic/gin"
	"tiktok/dao"
	"tiktok/setting"
)

func main() {
	//初始化setting
	setting.Init()
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
