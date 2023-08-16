package mysql_conf

import (
	"gopkg.in/ini.v1"
	"tiktok/lgr"
)

var (
	DSN      string
	UserName string
	Password string
	Host     string
	Schema   string
)

func init() {
	cfg, err := ini.Load("conf.ini")
	if err != nil {
		panic("mysql_conf ini文件读取异常")
	}
	UserName = cfg.Section("database").Key("username").String()
	Password = cfg.Section("database").Key("password").String()
	Host = cfg.Section("database").Key("host").String()
	Schema = cfg.Section("database").Key("schema").String()
	DSN = UserName + ":" + Password + "@tcp(" + Host + ")/" + Schema + "?charset=utf8mb4&parseTime=True"
	lgr.Print("Init mysql_conf " + DSN)
}
