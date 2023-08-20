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

// Create 保存视频信息
func (v *VideoDao) Create(video model.Video) (uint, error) {
	//操作Video表，保存video对象
	v.db.Create(&video)
	if err != nil {
		return
	}

}
