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
func (v VideoDao) CreateVideo(video model.Video) uint64 {
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
func (v VideoDao) ListVideoByAuthorId(authorId uint64) []model.Video {
	var videos []model.Video
	v.db.Where("author_id = ?", authorId).Find(&videos)
	return videos
}

// GetVideoById 通过视频id查找视频信息
func (v VideoDao) GetVideoById(videoId uint64) model.Video {
	video := model.Video{}
	v.db.Where("id = ?", videoId).Take(&video)
	return video
}

// CountVideoByAuthorId 通过作者的id计算视频数量
func (v VideoDao) CountVideoByAuthorId(authorId uint64) uint64 {
	ret := int64(0)
	v.db.Model(model.Video{}).Where("author_id = ?", authorId).Count(&ret)
	return uint64(ret)
}
