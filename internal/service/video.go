package service

import (
	"go.uber.org/zap"
	"mime/multipart"
	"tiktok/internal/conf"
	"tiktok/internal/data"
	"tiktok/internal/manager"
	"tiktok/internal/model"
	"tiktok/internal/terrs"
	"tiktok/pkg/util"
	"time"
)

type VideoService struct {
	logger          *zap.Logger
	ossUtil         *util.OssUtil
	voUtil          *util.VoUtil
	videoDao        *data.VideoDao
	userManager     *manager.UserManager
	favoriteManager *manager.FavoriteManager
	tiktokConf      *conf.TiktokConfig
}

func NewVideoService(
	zl *zap.Logger,
	ou *util.OssUtil,
	vu *util.VoUtil,
	vd *data.VideoDao,
	tc *conf.TiktokConfig,
	um *manager.UserManager,
	fm *manager.FavoriteManager) *VideoService {
	return &VideoService{
		logger:          zl,
		ossUtil:         ou,
		voUtil:          vu,
		videoDao:        vd,
		tiktokConf:      tc,
		userManager:     um,
		favoriteManager: fm,
	}
}

func (s VideoService) Publish(file *multipart.FileHeader, authorId uint64, title string) error {
	//使用对象存储工具类进行文件上传
	url, err := s.ossUtil.OSSUpload(file)
	if err != nil {
		s.logger.Error("文件上传出错", zap.String("cause", err.Error()))
		return terrs.ErrInternal
	}

	//没有出现错误，进行数据的储存。
	videoInfo := model.Video{
		Title:    title,
		AuthorId: authorId,
		PlayUrl:  url,
	}
	s.videoDao.CreateVideo(videoInfo)

	//没有错误，返回nil
	return nil
}

// Feed 返回符合要求的时间早于latestTime的视频列表以及该列表中时间最早的视频更新时间
func (s VideoService) Feed(latestTime time.Time) ([]model.Video, time.Time) {
	// 调用dao层代码查询视频信息
	videos := s.videoDao.ListVideoOrderByUpdateTime(s.tiktokConf.FeedSize, latestTime)
	if len(videos) == 0 {
		return []model.Video{}, time.Now()
	}
	//返回结果
	return videos, videos[len(videos)-1].UpdatedAt
}

// ListVideoByAuthorId 列出指定作者的所有作品
func (s VideoService) ListVideoByAuthorId(authorId uint64) []model.Video {
	//调用dao层代码查询视频信息
	videos := s.videoDao.ListVideoByAuthorId(authorId)

	//返回查询结果
	return videos
}
