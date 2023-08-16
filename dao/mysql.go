package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"tiktok/conf/mysql_conf"
	"tiktok/model/po"
)

var DB *gorm.DB

// InitMysql 初始化mysql连接
func init() {
	//初始化数据库连接
	DB, _ = gorm.Open(
		mysql.Open(mysql_conf.DSN), &gorm.Config{
			SkipDefaultTransaction: true, //关闭默认事务
			PrepareStmt:            true, //缓存预编译语句
		})
}

// InitTables 初始化数据表格
func InitTables() {
	err := DB.AutoMigrate(
		&po.User{}, &po.Video{}, &po.AuthorVideo{}, &po.Follow{},
		&po.Comment{}, &po.Favorite{}, &po.Message{},
	)
	if err != nil {
		fmt.Println("autoMigrate error !!!")
		return
	}
}
