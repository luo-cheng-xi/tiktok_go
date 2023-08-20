package conf

import (
	"gopkg.in/ini.v1"
)

type DefaultConfig struct {
	UserAvatar      string
	VideoTitle      string
	VideoCover      string
	BackGroundImage string
}

func GetDefaultConf() *DefaultConfig {
	cfg, err := ini.Load("./configs/conf.ini")
	if err != nil {
		panic("ini文件读取异常")
	}
	return &DefaultConfig{
		BackGroundImage: cfg.Section("default").Key("backgroundImage").String(),
		UserAvatar:      cfg.Section("default").Key("userAvatar").String(),
		VideoCover:      cfg.Section("default").Key("videoCover").String(),
		VideoTitle:      cfg.Section("default").Key("videoTitle").String(),
	}
}
