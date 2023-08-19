package util

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"go.uber.org/zap"
	"mime/multipart"
	"strings"
	"tiktok/internal/conf"
	"tiktok/pkg/logging"
)

var (
	_      *oss.Client
	Bucket *oss.Bucket
)

func loadOssConf() {
	Client, err := oss.New(conf.OSS.EndPoint, conf.OSS.AccessKeyId, conf.OSS.AccessKeySecret)
	if err != nil {
		panic("阿里云oss服务初始化异常")
	}
	Bucket, err = Client.Bucket(conf.OSS.BucketName)
	if err != nil {
		panic("阿里云ossBucket初始化异常")
	}
}

// init 完成阿里oss服务相关变量的初始化，因为只会使用一个bucket，所以变量中也只提供一个bucket了
func init() {
	Client, err := oss.New(conf.OSS.EndPoint, conf.OSS.AccessKeyId, conf.OSS.AccessKeySecret)
	if err != nil {
		panic("阿里云oss服务初始化异常")
	}
	Bucket, err = Client.Bucket(conf.OSS.BucketName)
	if err != nil {
		panic("阿里云ossBucket初始化异常")
	}
}

func OSSUpload(file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	path := file.Filename
	err = Bucket.PutObject(path, src)
	url := strings.Split(conf.OSS.EndPoint, "/")[0] + "/" + conf.OSS.BucketName + "." + strings.Split(conf.OSS.EndPoint, "/")[1] + path
	logging.Logger.Debug("上传文件至阿里云oss : ", zap.String("url", url))
	if err != nil {
		return "", err
	}
	return url, nil
}
