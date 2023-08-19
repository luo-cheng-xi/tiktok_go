package conf

import (
	"gopkg.in/ini.v1"
)

type ossConfig struct {
	EndPoint        string
	AccessKeyId     string
	AccessKeySecret string
	BucketName      string
}

var OSS = loadOSSConf()

func loadOSSConf() ossConfig {
	cfg, err := ini.Load("./configs/conf.ini")
	if err != nil {
		panic("oss_conf ini文件读取异常")
	}
	ret := ossConfig{
		EndPoint:        cfg.Section("oss").Key("endPoint").String(),
		AccessKeyId:     cfg.Section("oss").Key("accessKeyId").String(),
		AccessKeySecret: cfg.Section("oss").Key("accessKeySecret").String(),
		BucketName:      cfg.Section("oss").Key("bucketName").String(),
	}
	return ret
}