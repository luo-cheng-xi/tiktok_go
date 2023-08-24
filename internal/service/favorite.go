package service

import (
	"go.uber.org/zap"
	"tiktok/internal/manager"
	"tiktok/internal/model"
)

type FavoriteService struct {
	logger          *zap.Logger
	favoriteManager *manager.FavoriteManager
}

func NewFavoriteService(l *zap.Logger, fm *manager.FavoriteManager) *FavoriteService {
	return &FavoriteService{
		logger:          l,
		favoriteManager: fm,
	}
}

// FavoriteAction 登录用户对于视频的点赞和取消点赞操作
func (s FavoriteService) FavoriteAction(userId uint64, videoId uint64, actionType uint32) {
	if actionType == 1 {
		s.favoriteManager.SaveFavorite(userId, videoId)
	} else if actionType == 2 {
		s.favoriteManager.DeleteFavorite(userId, videoId)
	}
}

// ListFavoriteVideoByUserId 根据用户id列出所有该用户喜欢的视频
func (s FavoriteService) ListFavoriteVideoByUserId(userId uint64) []model.Video {
	//调用manager层代码获取目标用户喜欢的视频
	videos := s.favoriteManager.ListFavoriteVideoByUserId(userId)

	//返回结果
	return videos
}
