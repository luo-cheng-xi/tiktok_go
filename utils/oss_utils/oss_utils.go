package oss_utils

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"tiktok/conf/oss_conf"
)

var (
	Client *oss.Client
	Bucket *oss.Bucket
)

// init 完成阿里oss服务相关变量的初始化，因为只会使用一个bucket，所以变量中也只提供一个bucket了
func init() {
	Client, err := oss.New(oss_conf.EndPoint, oss_conf.AccessKeyId, oss_conf.AccessKeySecret)
	if err != nil {
		panic("阿里云oss服务初始化异常")
	}
	Bucket, err = Client.Bucket(oss_conf.BucketName)
	if err != nil {
		panic("阿里云ossBucket初始化异常")
	}
}
