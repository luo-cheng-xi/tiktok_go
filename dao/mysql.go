package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"tiktok/model/po"
	"tiktok/setting"
)

var (
	DB *gorm.DB
)

// InitMysql 初始化mysql连接
func InitMysql() {
	//dsn := "root:root123456@tcp(192.168.157.129:3306)/tiktok?charset=utf8mb4&parseTime=True"
	//初始化数据库连接
	DB, _ = gorm.Open(
		mysql.Open(setting.DatabaseDSN), &gorm.Config{
			SkipDefaultTransaction: true, //关闭默认事务
			PrepareStmt:            true, //缓存预编译语句
		})
}

// InitTables 初始化数据表格
func InitTables() {
	err := DB.AutoMigrate(
		&po.User{}, &po.Video{}, &po.AuthorVideo{}, &po.Follow{},
		&po.Comment{}, &po.Favorite{}, &po.Message{}, &po.Friend{},
	)
	if err != nil {
		fmt.Println("autoMigrate error !!!")
		return
	}
}
