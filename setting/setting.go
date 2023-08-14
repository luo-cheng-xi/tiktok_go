package setting

import (
	"gopkg.in/ini.v1"
	"log"
)

var (
	prefix           = "setting-----------------"
	DatabaseDSN      string
	DatabaseUserName string
	DatabasePassword string
	DatabaseHost     string
	DatabaseSchema   string

	JwtSignedKey string

	DefaultUserAvatar      string
	DefaultVideoCover      string
	DefaultBackGroundImage string
)

func Init() {
	logger := log.Default()
	cfg, err := ini.Load("conf.ini")
	if err != nil {
		return
	}
	DatabaseUserName = cfg.Section("database").Key("username").String()
	DatabasePassword = cfg.Section("database").Key("password").String()
	DatabaseHost = cfg.Section("database").Key("host").String()
	DatabaseSchema = cfg.Section("database").Key("schema").String()
	DatabaseDSN = DatabaseUserName + ":" + DatabasePassword + "@tcp(" + DatabaseHost + ")/" + DatabaseSchema + "?charset=utf8mb4&parseTime=True"
	logger.Println(prefix + "DatabaseDSN = " + DatabaseDSN)

	JwtSignedKey = cfg.Section("jwt").Key("signedKey").String()
	logger.Println(prefix + "SignedKey = " + JwtSignedKey)

	DefaultBackGroundImage = cfg.Section("default").Key("backgroundImage").String()
	DefaultUserAvatar = cfg.Section("default").Key("userAvatar").String()
	DefaultVideoCover = cfg.Section("default").Key("videoCover").String()
	logger.Println(prefix + "backgroundImage = " + DefaultBackGroundImage)
	logger.Println(prefix + "userAvatar = " + DefaultUserAvatar)
	logger.Println(prefix + "videoCover = " + DefaultVideoCover)
}
