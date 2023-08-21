package util

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/google/uuid"
	"mime/multipart"
	"path"
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

func (rx *OssUtil) GetUrl(fileName string) string {
	fmt.Println(rx.ossConf.EndPoint)
	parts := strings.Split(rx.ossConf.EndPoint, "//")
	return parts[0] + "//" + rx.ossConf.BucketName + "." + parts[1] + "/" + fileName
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
	ext := path.Ext(file.Filename)  //获取文件后缀名
	fileUUID := uuid.New().String() //生成文件uuid
	newName := fileUUID + ext
	err = Bucket.PutObject(newName, src)
	url := rx.GetUrl(newName)
	if err != nil {
		return "", err
	}
	return url, nil
}
