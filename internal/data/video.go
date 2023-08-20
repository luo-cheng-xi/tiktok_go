package data

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"tiktok/internal/model"
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
