package manager

import (
	"go.uber.org/zap"
	"tiktok/internal/data"
	"tiktok/internal/model"
	"tiktok/internal/terrs"
)

type RelationManager struct {
	logger      *zap.Logger
	relationDao *data.FollowDao
	userDao     *data.UserDao
}

func NewRelationManager(l *zap.Logger, rd *data.FollowDao, ud *data.UserDao) *RelationManager {
	return &RelationManager{
		logger:      l,
		relationDao: rd,
		userDao:     ud,
	}
}

// GetFollowCount 通过用户Id获得用户关注了的用户的人数
func (r RelationManager) GetFollowCount(userId uint64) uint64 {
	return r.relationDao.GetFollowCount(userId)
}

// GetFollowerCount 通过用户Id获得用户的粉丝数
func (r RelationManager) GetFollowerCount(userId uint64) uint64 {
	return r.relationDao.GetFollowerCount(userId)
}

// IsFollow 通过关注者和被关注者的id获取是否存在该关注关系
func (r RelationManager) IsFollow(userId uint64, toUserId uint64) (bool, error) {
	flag, err := r.relationDao.IsFollow(userId, toUserId)
	if err != nil {
		return false, err
	}
	return flag, nil
}

// SaveFollow 保存关注关系
func (r RelationManager) SaveFollow(userId, toUserId uint64) error {
	//检查是否已经关注了该用户，如果已经关注了，则返回用户已关注错误
	flag, err := r.relationDao.IsFollow(userId, toUserId)
	if err != nil {
		return err
	}
	if flag {
		return terrs.ErrUserFollowed
	}

	// 调用Dao层代码
	r.relationDao.SaveFollow(userId, toUserId)
	return nil
}

// DeleteFollow 删除关注关系
func (r RelationManager) DeleteFollow(userId, toUserid uint64) error {
	//检查是否关注了该用户，如果未关注，则返回用户未关注错误
	flag, err := r.relationDao.IsFollow(userId, toUserid)
	if err != nil {
		return err
	}
	if !flag {
		return terrs.ErrUserNotFollowed
	}
	r.relationDao.DeleteFollow(userId, toUserid)
	return nil
}

// ListFollowUser 列出指定id关注的所有用户的信息
func (r RelationManager) ListFollowUser(userId uint64) ([]model.User, error) {
	// 列出以该用户为关注者的所有follow关系
	follows := r.relationDao.ListFollowByFollowerId(userId)

	// 遍历follow关系,通过所有follow关系的followId查询user信息并封装
	users := make([]model.User, len(follows))
	for i, follow := range follows {
		var err error
		users[i], err = r.userDao.GetUserById(follow.FollowId)
		if err != nil {
			return nil, err
		}
	}
	return users, nil
}

// ListFollowerInfo 列出指定id所有粉丝信息
func (r RelationManager) ListFollowerInfo(userId uint64) ([]model.User, error) {
	// 列出一个该用户为被关注者的所有follow关系
	follows := r.relationDao.ListFollowByFollowId(userId)

	// 遍历follow关系,通过所有follow关系的followerId查询user信息并封装
	users := make([]model.User, len(follows))
	for i, follow := range follows {
		var err error
		users[i], err = r.userDao.GetUserById(follow.FollowerId)
		if err != nil {
			return nil, err
		}
	}
	return users, nil
}

// ListFriend 列出指定id的所有朋友信息
func (r RelationManager) ListFriend(userId uint64) ([]model.User, error) {
	//列出以该用户为关注者的所有信息
	asFollowers := r.relationDao.ListFollowByFollowerId(userId)
	//列出以该用户为被关注者的所有信息
	asFollows := r.relationDao.ListFollowByFollowId(userId)
	var friends []model.User
	for _, asFollower := range asFollowers {
		for _, asFollow := range asFollows {
			// 作为粉丝时的关注id,作为被关注者时的粉丝id,相等
			if asFollower.FollowId == asFollow.FollowerId {
				friendId := asFollower.FollowId
				info, err := r.userDao.GetUserById(friendId)
				if err != nil {
					r.logger.Debug("查询用户信息出错")
					return nil, err
				}
				friends = append(friends, info)
			}
		}
	}
	return friends, nil
}
