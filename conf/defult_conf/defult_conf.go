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

func Init(cfg *ini.File) {
	BackGroundImage = cfg.Section("default").Key("backgroundImage").String()
	UserAvatar = cfg.Section("default").Key("userAvatar").String()
	VideoCover = cfg.Section("default").Key("videoCover").String()
	lgr.Print("backgroundImage = " + BackGroundImage)
	lgr.Print("userAvatar = " + UserAvatar)
	lgr.Print("videoCover = " + VideoCover)
}
