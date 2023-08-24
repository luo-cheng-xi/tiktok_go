package service

import (
	"go.uber.org/zap"
	"tiktok/internal/manager"
	"tiktok/internal/model"
)

type CommentService struct {
	logger         *zap.Logger
	commentManager *manager.CommentManager
}

func NewCommentService(l *zap.Logger, c *manager.CommentManager) *CommentService {
	return &CommentService{
		logger:         l,
		commentManager: c,
	}
}

// CommentAction 执行评论行为，actionType = 1时发布，等于2时删除
func (s CommentService) CommentAction(userId uint64, videoId uint64, actionType uint32, commentText string, commentId uint64) model.Comment {
	if actionType == 1 {
		return s.commentManager.SaveComment(userId, videoId, commentText)
	} else if actionType == 2 {
		s.commentManager.DeleteComment(commentId)
		return model.Comment{}
	} else {
		return model.Comment{}
	}
}

// ListCommentByVideoId 基于视频id列出该视频的所有评论信息
func (s CommentService) ListCommentByVideoId(videoId uint64) []model.Comment {
	return s.commentManager.ListCommentByVideoId(videoId)
}
