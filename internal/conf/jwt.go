package conf

import (
	"gopkg.in/ini.v1"
)

type JwtConfig struct {
	SignedKey string
}

// GetJwtConf 获取Jwt配置
func GetJwtConf() *JwtConfig {
	cfg, err := ini.Load("./configs/conf.ini")
	if err != nil {
		panic("jwt_conf ini文件读取异常")
	}
	return &JwtConfig{
		SignedKey: cfg.Section("jwt").Key("signedKey").String(),
	}
}
