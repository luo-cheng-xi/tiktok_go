package service

import (
	"go.uber.org/zap"
	"mime/multipart"
	"tiktok/internal/conf"
	"tiktok/internal/data"
	"tiktok/internal/model"
	"tiktok/internal/terrs"
	"tiktok/pkg/util"
	"time"
)

type VideoService struct {
	logger     *zap.Logger
	ossUtil    *util.OssUtil
	videoDao   *data.VideoDao
	tiktokConf *conf.TiktokConfig
}

func NewVideoService(zl *zap.Logger, ou *util.OssUtil, dv *data.VideoDao, tc *conf.TiktokConfig) *VideoService {
	return &VideoService{
		logger:     zl,
		ossUtil:    ou,
		videoDao:   dv,
		tiktokConf: tc,
	}
}

func (s *VideoService) Publish(file *multipart.FileHeader, authorId uint, title string) error {
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
func (s *VideoService) Feed(latestTime time.Time) ([]model.VideoVO, time.Time) {
	// 调用dao层代码查询视频信息
	videos := s.videoDao.ListVideoOrderByUpdateTime(s.tiktokConf.FeedSize, latestTime)
	if len(videos) == 0 {
		return []model.VideoVO{}, time.Now()
	}
	//将实体类video,转化为前端所需数据videoVO。
	var videoVOs = make([]model.VideoVO, len(videos)) // 提前指定切片大小，避免动态扩容
	for i, video := range videos {
		videoVOs[i] = model.ParseVideoVO(video)
	}

	//返回结果
	return videoVOs, videos[len(videos)-1].UpdatedAt
}
