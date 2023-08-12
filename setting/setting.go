package setting

import "gopkg.in/ini.v1"

var (
	DatabaseDSN      string
	DatabaseUserName string
	DatabasePassword string
	DatabaseHost     string
	DatabaseSchema   string

	JwtSignedKey string
)

func Init() {
	cfg, err := ini.Load("conf.ini")
	if err != nil {
		return
	}
	DatabaseUserName = cfg.Section("database").Key("username").String()
	DatabasePassword = cfg.Section("database").Key("password").String()
	DatabaseHost = cfg.Section("database").Key("host").String()
	DatabaseSchema = cfg.Section("database").Key("schema").String()
	DatabaseDSN = DatabaseUserName + ":" + DatabasePassword + "@tcp(" + DatabaseHost + ")/" + DatabaseSchema + "?charset=utf8mb4&parseTime=True"

}
