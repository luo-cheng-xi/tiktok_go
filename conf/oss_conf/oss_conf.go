package oss_conf

import (
	"gopkg.in/ini.v1"
	"tiktok/lgr"
)

var (
	EndPoint        string
	AccessKeyId     string
	AccessKeySecret string
	BucketName      string
)

func init() {
	cfg, err := ini.Load("conf.ini")
	if err != nil {
		panic("oss_conf ini文件读取异常")
	}
	EndPoint = cfg.Section("oss").Key("endPoint").String()
	AccessKeyId = cfg.Section("oss").Key("accessKeyId").String()
	AccessKeySecret = cfg.Section("oss").Key("accessKeySecret").String()
	BucketName = cfg.Section("oss").Key("bucketName").String()
	lgr.Print("Init oss EndPoint = " + EndPoint)
	lgr.Print("Init oss AccessKeyId = " + AccessKeyId)
	lgr.Print("Init oss AccessKeySecrete = " + AccessKeySecret)
	lgr.Print("Init oss BucketName = " + BucketName)
}
