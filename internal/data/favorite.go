package data

import (
	"errors"
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

// SaveFavorite 在数据库Favorite表中添加点赞关系
func (f FavoriteDao) SaveFavorite(userid uint64, videoId uint64) {
	favorite := model.Favorite{
		UserId:  userid,
		VideoId: videoId,
	}
	//对于表中存在被逻辑删除字段的情况，更新deleted_at为null
	var ret model.Favorite
	if err := f.db.Unscoped().Where(favorite).Take(&ret).Error; err == nil {
		ret.DeletedAt = gorm.DeletedAt{}
		f.db.Save(ret)
	} else {
		f.db.Create(&favorite)
	}

}

// DeleteFavorite 在Favorite表中移除点赞关系
func (f FavoriteDao) DeleteFavorite(userId uint64, videoId uint64) {
	f.db.Where("user_id = ? and video_id = ?", userId, videoId).Delete(&model.Favorite{})
}

// ListFavoriteByUserId 通过用户Id查找与该用户关联的所有点赞关系
func (f FavoriteDao) ListFavoriteByUserId(userId uint64) []model.Favorite {
	var favorites []model.Favorite
	f.db.Where("user_id = ?", userId).Find(&favorites)
	return favorites
}

// CountFavoritedByVideoId 通过视频id查询获赞数量
func (f FavoriteDao) CountFavoritedByVideoId(videoId uint64) uint64 {
	ret := int64(0)
	f.db.Model(model.Favorite{}).Where("video_id = ?", videoId).Count(&ret)
	return uint64(ret)
}

// CountFavoriteByUserId 通过用户id查询他点赞了的视频
func (f FavoriteDao) CountFavoriteByUserId(userId uint64) uint64 {
	ret := int64(0)
	f.db.Model(model.Favorite{}).Where("user_id = ?", userId).Count(&ret)
	return uint64(ret)
}

// IsFavorite 查看用户是否点赞了该视频
func (f FavoriteDao) IsFavorite(userId, videoId uint64) (bool, error) {
	condition := model.Favorite{
		UserId:  userId,
		VideoId: videoId,
	}
	if err := f.db.Where(condition).Take(&model.Favorite{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		f.logger.Error("查询出错:", zap.String("cause", err.Error()))
		return false, err
	}
	return true, nil
}
