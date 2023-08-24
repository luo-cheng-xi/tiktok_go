package manager

import (
	"go.uber.org/zap"
	"tiktok/internal/data"
	"tiktok/internal/model"
)

type CommentManager struct {
	logger     *zap.Logger
	commentDao *data.CommentDao
}

func NewCommentManager(l *zap.Logger, cd *data.CommentDao) *CommentManager {
	return &CommentManager{
		logger:     l,
		commentDao: cd,
	}
}

// SaveComment 保存Comment，返回Comment信息
func (c CommentManager) SaveComment(userId, videoId uint64, content string) model.Comment {
	comment := model.Comment{
		UserId:  userId,
		VideoId: videoId,
		Content: content,
	}

	// 接收返回的comment,并返回
	return c.commentDao.SaveComment(comment)
}

// DeleteComment 删除Comment
func (c CommentManager) DeleteComment(id uint64) {
	c.commentDao.DeleteComment(id)
}

// ListCommentByVideoId 基于视频id列出所有评论信息
func (c CommentManager) ListCommentByVideoId(videoId uint64) []model.Comment {
	return c.commentDao.ListCommentByVideoId(videoId)
}

// CountCommentByVideoId 通过视频id得到视频评论数量
func (c CommentManager) CountCommentByVideoId(videoId uint64) uint64 {
	return c.commentDao.CountCommentByVideoId(videoId)
}
