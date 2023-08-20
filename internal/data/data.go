package data

import (
	"fmt"
	"github.com/google/wire"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"tiktok/internal/conf"
	"tiktok/internal/model"
	"tiktok/pkg/logging"
)

var ProviderSet = wire.NewSet(NewData, NewUserDao)

type Data struct {
	DB *gorm.DB
}

// 初始化数据库连接,设置dsn,设置gorm日志为自定义的zap日志
func newConn(d *conf.Data, logger gormlogger.Interface) (*gorm.DB, error) {
	var dsn string
	dsn = d.Database.DSN
	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true, //关闭默认事务
		PrepareStmt:            true, //缓存预编译语句
		Logger:                 logger,
	})
}

// autoMigrate 迁移数据库
func autoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&model.User{},
		&model.Video{},
		//&model.AuthorVideo{},
		&model.Follow{},
		&model.Comment{},
		&model.Favorite{},
		&model.Message{},
	)
	if err != nil {
		fmt.Println("autoMigrate error!!!")
		return err
	}
	return nil
}

// NewData 初始化数据集合
func NewData(d *conf.Data, l *zap.Logger) (*Data, error) {
	gormLogger := logging.NewGormLogger(l)
	conn, err := newConn(d, gormLogger)
	if err != nil {
		return nil, err
	}
	err = autoMigrate(conn)
	if err != nil {
		return nil, err
	}
	return &Data{
		DB: conn,
	}, nil
}
