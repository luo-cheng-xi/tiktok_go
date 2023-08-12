package main

import (
	"github.com/gin-gonic/gin"
	"tiktok/dao"
	"tiktok/setting"
)

type User struct {
	ID              uint
	Name            string
	Password        string
	Avatar          string
	BackgroundImage string
	Signature       string
	IsDelete        bool
}

func (User) TableName() string {
	return "user"
}

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
