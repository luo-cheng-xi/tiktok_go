package jwt_conf

import (
	"gopkg.in/ini.v1"
	"tiktok/lgr"
)

var (
	JwtSignedKey string
)

func init() {
	cfg, err := ini.Load("conf.ini")
	if err != nil {
		panic("jwt_conf ini文件读取异常")
	}
	JwtSignedKey = cfg.Section("jwt").Key("signedKey").String()
	lgr.Print("Init SignedKey = " + JwtSignedKey)
}
