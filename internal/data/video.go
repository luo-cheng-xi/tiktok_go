package data

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"tiktok/internal/model"
	"time"
)

type VideoDao struct {
	logger *zap.Logger
	db     *gorm.DB
}

func NewVideoDao(l *zap.Logger, d *Data) *VideoDao {
	return &VideoDao{
		logger: l,
		db:     d.DB,
	}
}

// CreateVideo 保存视频信息,返回视频id
func (v VideoDao) CreateVideo(video model.Video) int64 {
	//操作Video表，保存video对象
	v.db.Create(&video)
	return video.ID
}

// ListVideoOrderByUpdateTime 根据更新时间倒序排序，传入查询大小限制
func (v VideoDao) ListVideoOrderByUpdateTime(limit int, latestTime time.Time) []model.Video {
	var videos []model.Video
	v.db.Where("updated_at < ?", latestTime).Order("updated_at desc").Limit(limit).Find(&videos)
	return videos
}

// ListVideoByAuthorId 列出指定作者Id的所有视频
func (v VideoDao) ListVideoByAuthorId(authorId int64) []model.Video {
	var videos []model.Video
	v.db.Where("author_id = ?", authorId).Find(&videos)
	return videos
}

// Favorite 在数据库Favorite表中添加点赞关系
func (v VideoDao) Favorite(userid int64, videoId int64) {
	favorite := model.Favorite{
		UserId:  userid,
		VideoId: videoId,
	}
	v.db.Create(favorite)
}

// CancelFavorite 在Favorite表中移除点赞关系
func (v VideoDao) CancelFavorite(userId int64, videoId int64) {
	v.db.Where("where user_id = ? and video_id = ?", userId, videoId).Delete(&model.Favorite{})
}
