package conf

import (
	"gopkg.in/ini.v1"
)

// 默认值配置
type defaultConfig struct {
	UserAvatar      string
	VideoTitle      string
	VideoCover      string
	BackGroundImage string
}

var Default = loadDefaultConf()

func loadDefaultConf() defaultConfig {
	cfg, err := ini.Load("./configs/conf.ini")
	if err != nil {
		panic("ini文件读取异常")
	}
	ret := defaultConfig{
		BackGroundImage: cfg.Section("default").Key("backgroundImage").String(),
		UserAvatar:      cfg.Section("default").Key("userAvatar").String(),
		VideoCover:      cfg.Section("default").Key("videoCover").String(),
		VideoTitle:      cfg.Section("default").Key("videoTitle").String(),
	}
	return ret
}
