package util

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"mime/multipart"
	"strings"
	"tiktok/internal/conf"
)

var (
	_      *oss.Client
	Bucket *oss.Bucket
)

type OssUtil struct {
	ossConf *conf.OssConfig
}

func GetOssUtil(oc *conf.OssConfig) *OssUtil {
	return &OssUtil{
		ossConf: oc,
	}
}
func (rx *OssUtil) loadOssConf() {
	Client, err := oss.New(rx.ossConf.EndPoint, rx.ossConf.AccessKeyId, rx.ossConf.AccessKeySecret)
	if err != nil {
		panic("阿里云oss服务初始化异常")
	}
	Bucket, err = Client.Bucket(rx.ossConf.BucketName)
	if err != nil {
		panic("阿里云ossBucket初始化异常")
	}
}

// init 完成阿里oss服务相关变量的初始化，因为只会使用一个bucket，所以变量中也只提供一个bucket了
func (rx *OssUtil) init() {
	Client, err := oss.New(rx.ossConf.EndPoint, rx.ossConf.AccessKeyId, rx.ossConf.AccessKeySecret)
	if err != nil {
		panic("阿里云oss服务初始化异常")
	}
	Bucket, err = Client.Bucket(rx.ossConf.BucketName)
	if err != nil {
		panic("阿里云ossBucket初始化异常")
	}
}

func (rx *OssUtil) OSSUpload(file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	path := file.Filename
	err = Bucket.PutObject(path, src)
	url := strings.Split(rx.ossConf.EndPoint, "/")[0] + "/" + rx.ossConf.BucketName + "." + strings.Split(rx.ossConf.EndPoint, "/")[1] + path
	if err != nil {
		return "", err
	}
	return url, nil
}
