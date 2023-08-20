package service

import (
	"go.uber.org/zap"
	"mime/multipart"
	"tiktok/internal/terrs"
	"tiktok/pkg/util"
)

type VideoService struct {
	logger  *zap.Logger
	ossUtil *util.OssUtil
}

func NewVideoService(l *zap.Logger, o *util.OssUtil) *VideoService {
	return &VideoService{
		logger:  l,
		ossUtil: o,
	}
}

func (s VideoService) Publish(file *multipart.FileHeader) error {
	url, err := s.ossUtil.OSSUpload(file)
	if err != nil {
		s.logger.Error("阿里oss文件上传出错",
			zap.String("cause", err.Error()))
		return terrs.ErrInternal
	}
	//没有出现错误，进行数据的储存。

}
