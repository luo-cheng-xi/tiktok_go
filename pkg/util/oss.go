package util

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"mime/multipart"
	"strings"
	"tiktok/internal/conf"
)

type OssUtil struct {
	ossConf *conf.OssConfig
}

func GetOssUtil(oc *conf.OssConfig) *OssUtil {
	return &OssUtil{
		ossConf: oc,
	}
}

// OSSUpload 上传文件至阿里云oss对象存储服务
func (rx *OssUtil) OSSUpload(file *multipart.FileHeader) (string, error) {
	Client, err := oss.New(rx.ossConf.EndPoint, rx.ossConf.AccessKeyId, rx.ossConf.AccessKeySecret)
	if err != nil {
		panic("阿里云oss服务初始化异常")
	}
	Bucket, err := Client.Bucket(rx.ossConf.BucketName)
	if err != nil {
		panic("阿里云ossBucket初始化异常")
	}
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
