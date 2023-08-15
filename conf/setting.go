package conf

import (
	"gopkg.in/ini.v1"
	"tiktok/conf/defult_conf"
	"tiktok/conf/jwt_conf"
	"tiktok/conf/mysql_conf"
	"tiktok/conf/redis_conf"
)

func Init() {
	cfg, err := ini.Load("conf.ini")
	if err != nil {
		return
	}
	jwt_conf.Init(cfg)
	mysql_conf.Init(cfg)
	redis_conf.Init(cfg)
	defult_conf.Init(cfg)
}
