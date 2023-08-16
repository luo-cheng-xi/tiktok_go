package defult_conf

import (
	"gopkg.in/ini.v1"
	"tiktok/lgr"
)

var (
	UserAvatar      string
	VideoCover      string
	BackGroundImage string
)

func init() {
	cfg, err := ini.Load("conf.ini")
	if err != nil {
		panic("ini文件读取异常")
	}
	BackGroundImage = cfg.Section("default").Key("backgroundImage").String()
	UserAvatar = cfg.Section("default").Key("userAvatar").String()
	VideoCover = cfg.Section("default").Key("videoCover").String()
	lgr.Print("Init backgroundImage = " + BackGroundImage)
	lgr.Print("Init userAvatar = " + UserAvatar)
	lgr.Print("Init videoCover = " + VideoCover)
}
