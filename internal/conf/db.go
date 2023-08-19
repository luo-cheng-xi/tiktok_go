package conf

import (
	"gopkg.in/ini.v1"
)

type dbConfig struct {
	DSN      string
	UserName string
	Password string
	Host     string
	Schema   string
}

var DB = loadMysqlConfig()

func loadMysqlConfig() dbConfig {
	cfg, err := ini.Load("./configs/conf.ini")
	if err != nil {
		panic("mysql_conf ini文件读取异常")
	}
	userName := cfg.Section("database").Key("username").String()
	password := cfg.Section("database").Key("password").String()
	host := cfg.Section("database").Key("host").String()
	schema := cfg.Section("database").Key("schema").String()
	dsn := userName + ":" + password + "@tcp(" + host + ")/" + schema + "?charset=utf8mb4&parseTime=True"
	ret := dbConfig{
		UserName: userName,
		Password: password,
		Host:     host,
		Schema:   schema,
		DSN:      dsn,
	}
	return ret
}
