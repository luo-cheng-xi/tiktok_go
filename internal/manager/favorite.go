package manager

import (
	"go.uber.org/zap"
	"tiktok/internal/data"
	"tiktok/internal/model"
)

type FavoriteManager struct {
	logger      *zap.Logger
	userDao     *data.UserDao
	videoDao    *data.VideoDao
	favoriteDao *data.FavoriteDao
}

func NewFavoriteManager(l *zap.Logger, d *data.UserDao, fd *data.FavoriteDao, vd *data.VideoDao) *FavoriteManager {
	return &FavoriteManager{
		logger:      l,
		userDao:     d,
		favoriteDao: fd,
		videoDao:    vd,
	}
}

// GetTotalFavorited 计算用户的被赞次数
// 逻辑上联表查 video-favorite
func (f FavoriteManager) GetTotalFavorited(userId uint64) uint64 {
	//列出该用户的所有作品
	videos := f.videoDao.ListVideoByAuthorId(userId)

	total := uint64(0)
	//遍历所有作品，记录这些作品的获赞数
	for _, video := range videos {
		total += f.favoriteDao.CountFavoritedByVideoId(video.ID)
	}

	//返回结果
	return total
}

// ListFavoriteVideoByUserId 根据用户id列出用户喜欢的视频
func (f FavoriteManager) ListFavoriteVideoByUserId(userId uint64) []model.Video {
	//调用dao层代码，查询目标用户相关的所有favorite关系
	favorites := f.favoriteDao.ListFavoriteByUserId(userId)

	//根据favorite信息查询所有与之相关的视频
	var videos = make([]model.Video, len(favorites))
	for i, fav := range favorites {
		videos[i] = f.videoDao.GetVideoById(fav.VideoId)
	}

	//返回结果
	return videos
}

// CountFavoriteByUserId 通过用户Id信息得到用户喜欢的视频数
func (f FavoriteManager) CountFavoriteByUserId(userId uint64) uint64 {
	return f.favoriteDao.CountFavoriteByUserId(userId)
}

// CountFavoritedByVideoId 通过视频Id获得视频的获赞数
func (f FavoriteManager) CountFavoritedByVideoId(videoId uint64) uint64 {
	return f.favoriteDao.CountFavoritedByVideoId(videoId)
}

// IsFavorite 通过用户和视频id查看用户是否点赞了该视频
func (f FavoriteManager) IsFavorite(userId, videoId uint64) (bool, error) {
	return f.favoriteDao.IsFavorite(userId, videoId)
}

// SaveFavorite 添加点赞关系
func (f FavoriteManager) SaveFavorite(userId uint64, videoId uint64) {
	f.favoriteDao.SaveFavorite(model.Favorite{UserId: userId, VideoId: videoId})
}

// DeleteFavorite 取消点赞关系
func (f FavoriteManager) DeleteFavorite(userId uint64, videoId uint64) {
	f.favoriteDao.DeleteFavorite(userId, videoId)
}
