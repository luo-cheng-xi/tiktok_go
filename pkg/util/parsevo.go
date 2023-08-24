package util

import (
	"go.uber.org/zap"
	"tiktok/internal/manager"
	"tiktok/internal/model"
)

type VoUtil struct {
	logger          *zap.Logger
	userManager     *manager.UserManager
	videoManager    *manager.VideoManager
	relationManager *manager.RelationManager
	favoriteManager *manager.FavoriteManager
	commentManager  *manager.CommentManager
}

func NewVoUtil(
	l *zap.Logger,
	um *manager.UserManager,
	vm *manager.VideoManager,
	rm *manager.RelationManager,
	fm *manager.FavoriteManager,
	cm *manager.CommentManager) *VoUtil {
	return &VoUtil{
		logger:          l,
		userManager:     um,
		videoManager:    vm,
		relationManager: rm,
		favoriteManager: fm,
		commentManager:  cm,
	}
}

// ParseUserVO 将user转化为UserVO
func (v VoUtil) ParseUserVO(tarUser model.User, curUserId uint64) (model.UserVO, error) {
	isFollow, err := v.relationManager.IsFollow(curUserId, tarUser.ID)
	if err != nil {
		return model.UserVO{}, err
	}
	return model.UserVO{
		ID:              tarUser.ID,
		Name:            tarUser.Username,
		FollowCount:     v.relationManager.GetFollowCount(tarUser.ID),
		FollowerCount:   v.relationManager.GetFollowerCount(tarUser.ID),
		IsFollow:        isFollow,
		Avatar:          tarUser.Avatar,
		BackgroundImage: tarUser.BackgroundImage,
		Signature:       tarUser.Signature,
		TotalFavorited:  v.favoriteManager.GetTotalFavorited(tarUser.ID),
		WorkCount:       v.videoManager.CountVideoByAuthorId(tarUser.ID),
		FavoriteCount:   v.favoriteManager.CountFavoriteByUserId(tarUser.ID),
	}, nil
}

// ParseVideoVO 待补全
func (v VoUtil) ParseVideoVO(video model.Video, curUserId uint64) (model.VideoVO, error) {
	author, err := v.userManager.GetUserById(video.AuthorId)
	if err != nil {
		return model.VideoVO{}, err
	}
	authorVO, err := v.ParseUserVO(author, curUserId)
	if err != nil {
		return model.VideoVO{}, err
	}
	isFavorite, err := v.favoriteManager.IsFavorite(curUserId, video.ID)
	if err != nil {
		return model.VideoVO{}, err
	}
	return model.VideoVO{
		ID:            video.ID,
		Author:        authorVO,
		PlayUrl:       video.PlayUrl,
		CoverUrl:      video.CoverUrl,
		FavoriteCount: v.favoriteManager.CountFavoritedByVideoId(video.ID),
		CommentCount:  v.commentManager.CountCommentByVideoId(video.ID),
		IsFavorite:    isFavorite,
		Title:         video.Title,
	}, nil
}

// ParseCommentVO 将comment转化为commentVO
func (v VoUtil) ParseCommentVO(comment model.Comment, curUserId uint64) (model.CommentVO, error) {
	user, err := v.userManager.GetUserById(comment.UserId)
	if err != nil {
		return model.CommentVO{}, err
	}
	userVO, err := v.ParseUserVO(user, curUserId)
	if err != nil {
		return model.CommentVO{}, err
	}
	return model.CommentVO{
		ID:         comment.ID,
		User:       userVO,
		Content:    comment.Content,
		CreateData: comment.CreatedAt.Format("01-02"),
	}, nil
}
