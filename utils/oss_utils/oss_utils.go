package oss_utils

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"tiktok/conf/oss_conf"
)

var (
	client *oss.Client
	bucket *oss.Bucket
)

func init() {
	client, err := oss.New(oss_conf.EndPoint, oss_conf.AccessKeyId, oss_conf.AccessKeySecret)
	if err != nil {
		panic("阿里云oss服务初始化异常")
	}
	bucket, err = client.Bucket(oss_conf.BucketName)
	if err != nil {
		panic("阿里云ossBucket初始化异常")
	}
}
