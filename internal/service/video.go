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
	logger      *zap.Logger
	ossUtil     *util.OssUtil
	videoDao    *data.VideoDao
	favoriteDao *data.FavoriteDao
	tiktokConf  *conf.TiktokConfig
}

func NewVideoService(zl *zap.Logger, ou *util.OssUtil, vd *data.VideoDao, tc *conf.TiktokConfig, fd *data.FavoriteDao) *VideoService {
	return &VideoService{
		logger:      zl,
		ossUtil:     ou,
		videoDao:    vd,
		favoriteDao: fd,
		tiktokConf:  tc,
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
func (s VideoService) Feed(latestTime time.Time) ([]model.VideoVO, time.Time) {
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

// ListVideoByAuthorId 列出指定作者的所有作品
func (s VideoService) ListVideoByAuthorId(authorId uint64) []model.VideoVO {
	//调用dao层代码查询视频信息
	videos := s.videoDao.ListVideoByAuthorId(authorId)

	//将实体类转化为前端所需数据videoVO
	var videoVOs = make([]model.VideoVO, len(videos))
	for i, video := range videos {
		videoVOs[i] = model.ParseVideoVO(video)
	}
	//返回结果
	return videoVOs
}

// FavoriteAction 登录用户对于视频的点赞和取消点赞操作
func (s VideoService) FavoriteAction(userId uint64, videoId uint64, actionType uint32) {
	if actionType == 1 {
		s.favoriteDao.Favorite(userId, videoId)
	} else if actionType == 2 {
		s.favoriteDao.CancelFavorite(userId, videoId)
	}
}

// ListFavoriteVideoByUserId 根据用户id列出所有该用户喜欢的视频
func (s VideoService) ListFavoriteVideoByUserId(curUserId uint64, tarUserId uint64) []model.VideoVO {
	//调用dao层代码，查询目标用户相关的所有favorite关系
	favorites := s.favoriteDao.ListFavoriteByUserId(tarUserId)

	//将视频信息都转化为videoVO
	var videoVOs = make([]model.VideoVO, len(favorites))
	for i, fav := range favorites {
		video := s.videoDao.GetVideoById(fav.VideoId)
		videoVOs[i] = model.ParseVideoVO(video)
	}

	//返回结果
	return videoVOs
}
