package data

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"tiktok/internal/model"
)

type CommentDao struct {
	logger *zap.Logger
	db     *gorm.DB
}

func NewCommentDao(l *zap.Logger, data *Data) *CommentDao {
	return &CommentDao{
		logger: l,
		db:     data.DB,
	}
}

// SaveComment 保存评论信息,返回评论信息
func (r CommentDao) SaveComment(comment model.Comment) model.Comment {
	r.db.Create(&comment)
	return comment
}

// DeleteComment 删除评论信息
func (r CommentDao) DeleteComment(id uint64) {
	r.db.Delete(&model.Comment{}, id)
}

// ListCommentByVideoId 查找videoId为指定id的评论信息
func (r CommentDao) ListCommentByVideoId(videoId uint64) []model.Comment {
	var comments []model.Comment
	r.db.Where("video_id = ?", videoId).Find(&comments)
	return comments
}

// CountCommentByVideoId 通过视频id查找评论数量
func (r CommentDao) CountCommentByVideoId(videoId uint64) uint64 {
	ret := int64(0)
	r.db.Model(&model.Comment{}).Where("video_id = ?", videoId).Count(&ret)
	return uint64(ret)
}
