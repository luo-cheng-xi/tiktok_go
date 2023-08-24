package data

import (
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"tiktok/internal/model"
)

type FollowDao struct {
	logger *zap.Logger
	db     *gorm.DB
}

func NewRelationDao(zl *zap.Logger, data *Data) *FollowDao {
	return &FollowDao{
		logger: zl,
		db:     data.DB,
	}
}

// GetFollowCount 根据用户的ID获取用户关注的用户数
func (r FollowDao) GetFollowCount(followerId uint64) uint64 {
	var ret int64
	r.db.Model(model.Follow{}).Where("follower_id = ?", followerId).Count(&ret)
	return uint64(ret)
}

// GetFollowerCount 根据用户Id获取用户粉丝信息
func (r FollowDao) GetFollowerCount(followId uint64) uint64 {
	var ret int64
	r.db.Model(model.Follow{}).Where("follow_id = ?", followId).Count(&ret)
	return uint64(ret)
}

// SaveFollow 记录用户关注关系
func (r FollowDao) SaveFollow(followerId uint64, followId uint64) {
	follow := model.Follow{
		FollowerId: followerId,
		FollowId:   followId,
	}
	// 对于表中存在被逻辑删除了的字段的情况，更新deleted_at为null
	var ret model.Follow
	if err := r.db.Unscoped().Where(follow).Take(&ret).Error; err == nil {
		ret.DeletedAt = gorm.DeletedAt{}
		r.db.Save(ret)
	} else {
		r.db.Save(&follow)
	}
}

// DeleteFollow 删除用户关注关系
func (r FollowDao) DeleteFollow(followerId uint64, followId uint64) {
	r.db.Where("follower_id = ? and follow_id = ?", followerId, followId).Delete(&model.Follow{})
}

// IsFollow 查询是否存在指定的用户关注关系
// 如果存在返回真，不存在返回假
func (r FollowDao) IsFollow(userId uint64, toUserId uint64) (bool, error) {
	follow := model.Follow{}
	err := r.db.Where(model.Follow{FollowerId: userId, FollowId: toUserId}).Take(&follow).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		r.logger.Error("IsFollow查询出错", zap.String("cause", err.Error()))
		return false, err
	}
	return true, nil
}

// ListFollowByFollowerId 根据关注者id列出关注关系
func (r FollowDao) ListFollowByFollowerId(followerId uint64) []model.Follow {
	var follows []model.Follow
	r.db.Where("follower_id = ?", followerId).Find(&follows)
	return follows
}

// ListFollowByFollowId 根据被关注者id列出关注关系
func (r FollowDao) ListFollowByFollowId(followId uint64) []model.Follow {
	var follows []model.Follow
	r.db.Where("follow_id = ?", followId).Find(&follows)
	return follows
}
