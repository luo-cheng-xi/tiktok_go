package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"tiktok/model"
	"tiktok/setting"
)

var (
	DB *gorm.DB
)

func InitMysql() {
	//dsn := "root:root123456@tcp(192.168.157.129:3306)/tiktok?charset=utf8mb4&parseTime=True"
	//初始化数据库连接
	DB, _ = gorm.Open(
		mysql.Open(setting.DatabaseDSN), &gorm.Config{
			SkipDefaultTransaction: true, //关闭默认事务
			PrepareStmt:            true, //缓存预编译语句
		})
}

func InitTables() {
	err := DB.AutoMigrate(
		&model.User{}, &model.Video{}, &model.AuthorVideo{}, &model.Follow{},
		&model.Comment{}, &model.Favorite{}, &model.Message{}, &model.Friend{},
	)
	if err != nil {
		fmt.Println("autoMigrate error !!!")
		return
	}
}
