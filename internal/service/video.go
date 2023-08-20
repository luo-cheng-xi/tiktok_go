package service

import (
	"go.uber.org/zap"
	"mime/multipart"
	"tiktok/internal/data"
	"tiktok/internal/model"
	"tiktok/internal/terrs"
	"tiktok/pkg/util"
)

type VideoService struct {
	logger   *zap.Logger
	ossUtil  *util.OssUtil
	videoDao *data.VideoDao
}

func NewVideoService(zl *zap.Logger, ou *util.OssUtil, dv *data.VideoDao) *VideoService {
	return &VideoService{
		logger:   zl,
		ossUtil:  ou,
		videoDao: dv,
	}
}

func (s VideoService) Publish(file *multipart.FileHeader, authorId uint) error {
	//使用对象存储工具类进行文件上传
	url, err := s.ossUtil.OSSUpload(file)
	if err != nil {
		s.logger.Error("文件上传出错", zap.String("cause", err.Error()))
		return terrs.ErrInternal
	}

	//没有出现错误，进行数据的储存。
	videoInfo := model.Video{
		AuthorId: authorId,
		PlayUrl:  url,
	}
	s.videoDao.CreateVideo(videoInfo)

	//没有错误，返回nil
	return nil
}
