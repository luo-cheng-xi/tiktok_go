package data

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"tiktok/internal/model"
)

type FavoriteDao struct {
	logger *zap.Logger
	db     *gorm.DB
}

func NewFavoriteDao(l *zap.Logger, data *Data) *FavoriteDao {
	return &FavoriteDao{
		logger: l,
		db:     data.DB,
	}
}

// Favorite 在数据库Favorite表中添加点赞关系
func (f FavoriteDao) Favorite(userid uint64, videoId uint64) {
	favorite := model.Favorite{
		UserId:  userid,
		VideoId: videoId,
	}
	f.db.Create(favorite)
}

// CancelFavorite 在Favorite表中移除点赞关系
func (f FavoriteDao) CancelFavorite(userId uint64, videoId uint64) {
	f.db.Where("where user_id = ? and video_id = ?", userId, videoId).Delete(&model.Favorite{})
}

// ListFavoriteByUserId 通过用户Id查找与该用户关联的所有点赞关系
func (f FavoriteDao) ListFavoriteByUserId(userId uint64) []model.Favorite {
	var favorites []model.Favorite
	f.db.Where("user_id = ?", userId).Find(&favorites)
	return favorites
}
