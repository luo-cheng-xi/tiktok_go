package conf

import (
	"gopkg.in/ini.v1"
	"log"
)

type TiktokConfig struct {
	// DefaultUserAvatar 默认用户头像
	DefaultUserAvatar string
	// DefaultVideoTitle 默认标题
	DefaultVideoTitle string
	// DefaultVideoCover 默认视频封面
	DefaultVideoCover string
	// DefaultBackGroundImage 默认用户背景
	DefaultBackGroundImage string
	// FeedSize 视频流最大容量
	FeedSize int
}

func GetTiktokConf() *TiktokConfig {
	cfg, err := ini.Load(getIniPath())
	if err != nil {
		panic("ini文件读取异常")
	}
	return &TiktokConfig{
		DefaultBackGroundImage: cfg.Section("tiktok").Key("backgroundImage").String(), // backgroundImage
		DefaultUserAvatar:      cfg.Section("tiktok").Key("userAvatar").String(),      // userAvatar
		DefaultVideoCover:      cfg.Section("tiktok").Key("videoCover").String(),      // videoCover
		DefaultVideoTitle:      cfg.Section("tiktok").Key("videoTitle").String(),      // videoTitle
		FeedSize: func() int {
			size, err := cfg.Section("tiktok").Key("feedSize").Int()
			if err != nil {
				log.Fatal("视频流容量初始化错误")
			}
			return size
		}(), // feedSize
	}
}
