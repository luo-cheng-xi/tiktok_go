package conf

import (
	"gopkg.in/ini.v1"
)

type jwtConfig struct {
	SignedKey string
}

var Jwt = loadJwtConfig()

func loadJwtConfig() jwtConfig {
	cfg, err := ini.Load("./configs/conf.ini")
	if err != nil {
		panic("jwt_conf ini文件读取异常")
	}
	ret := jwtConfig{
		cfg.Section("jwt").Key("signedKey").String(),
	}
	return ret
}
