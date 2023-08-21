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
func (v *VideoDao) CreateVideo(video model.Video) uint {
	//操作Video表，保存video对象
	v.db.Create(&video)
	return video.ID
}

// ListVideoOrderByUpdateTime 根据更新时间倒序排序，传入查询大小限制
func (v *VideoDao) ListVideoOrderByUpdateTime(limit int, latestTime time.Time) []model.Video {
	var videos []model.Video
	v.db.Where("updated_at < ?", latestTime).Order("updated_at desc").Limit(limit).Find(&videos)
	return videos
}
