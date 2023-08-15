package jwt_conf

import (
	"gopkg.in/ini.v1"
	"tiktok/lgr"
)

var (
	JwtSignedKey string
)

func Init(cfg *ini.File) {
	JwtSignedKey = cfg.Section("jwt").Key("signedKey").String()
	lgr.Print("SignedKey = " + JwtSignedKey)
}
