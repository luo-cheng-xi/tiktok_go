package data

import (
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"tiktok/internal/model"
)

type RelationDao struct {
	logger *zap.Logger
	db     *gorm.DB
}

func NewRelationDao(zl *zap.Logger, data *Data) *RelationDao {
	return &RelationDao{
		logger: zl,
		db:     data.DB,
	}
}

// GetFollowCount 根据用户的ID获取用户关注的用户数
func (r RelationDao) GetFollowCount(followerId uint64) uint64 {
	var ret int64
	r.db.Model(model.Follow{}).Where("follower_id = ?", followerId).Count(&ret)
	return uint64(ret)
}

// GetFollowerCount 根据用户Id获取用户粉丝信息
func (r RelationDao) GetFollowerCount(followId uint64) uint64 {
	var ret int64
	r.db.Model(model.Follow{}).Where("follow_id = ?", followId).Count(&ret)
	return uint64(ret)
}

// Follow 记录用户关注关系
func (r RelationDao) Follow(userId uint64, toUserId uint64) {
	follow := model.Follow{
		FollowerId: userId,
		FollowId:   toUserId,
	}
	r.db.Create(&follow)
}

// IsFollow 查询是否存在指定的用户关注关系
// 如果存在返回真，不存在返回假
func (r RelationDao) IsFollow(userId uint64, toUserId uint64) (bool, error) {
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
