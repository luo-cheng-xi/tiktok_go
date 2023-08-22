package manager

import (
	"go.uber.org/zap"
	"tiktok/internal/data"
)

type Manager struct {
	logger      *zap.Logger
	userDao     *data.UserDao
	videoDao    *data.VideoDao
	favoriteDao *data.FavoriteDao
}

func NewManager(l *zap.Logger, d *data.UserDao) *Manager {
	return &Manager{
		logger:  l,
		userDao: d,
	}
}

// GetTotalFavorited 计算用户的被赞次数
// 逻辑上联表查 video-favorite
func (m *Manager) GetTotalFavorited(userId uint64) uint64 {
	//列出该用户的所有作品
	videos := m.videoDao.ListVideoByAuthorId(userId)

	total := uint64(0)
	//遍历所有作品，记录这些作品的获赞数
	for _, video := range videos {
		total += m.favoriteDao.CountFavoritedByVideoId(video.ID)
	}

	//返回结果
	return total
}
