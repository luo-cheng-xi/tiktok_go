package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"tiktok/internal/conf"
	"tiktok/internal/model"
)

var DB = loadDB()

func loadDB() *gorm.DB {
	//初始化数据库连接
	DB, err := gorm.Open(
		mysql.Open(conf.DB.DSN), &gorm.Config{
			SkipDefaultTransaction: true, //关闭默认事务
			PrepareStmt:            true, //缓存预编译语句
		})
	if err != nil {
		panic(err)
	}
	return DB
}

// InitTables 初始化数据表格
func InitTables() {
	err := DB.AutoMigrate(
		&model.User{}, &model.Video{}, &model.AuthorVideo{}, &model.Follow{},
		&model.Comment{}, &model.Favorite{}, &model.Message{},
	)
	if err != nil {
		fmt.Println("autoMigrate error!!!")
		return
	}
}
