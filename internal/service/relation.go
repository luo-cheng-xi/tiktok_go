package service

import (
	"go.uber.org/zap"
	"tiktok/internal/data"
	"tiktok/internal/manager"
	"tiktok/internal/model"
	"tiktok/internal/terrs"
)

type RelationService struct {
	logger          *zap.Logger
	relationManager *manager.RelationManager
	relationDao     *data.FollowDao
}

func NewRelationService(zl *zap.Logger, rm *manager.RelationManager, rd *data.FollowDao) *RelationService {
	return &RelationService{
		logger:          zl,
		relationManager: rm,
		relationDao:     rd,
	}
}

func (r RelationService) FollowAction(userId, toUserId, actionType uint64) error {
	switch actionType {
	case 1: //如果actionType为1，保存关注关系
		err := r.relationManager.SaveFollow(userId, toUserId)
		if err != nil {
			return err
		}
	case 2: //如果actionType为2，删除关注关系
		err := r.relationManager.DeleteFollow(userId, toUserId)
		if err != nil {
			return err
		}
	default:
		return terrs.ErrParamInvalid
	}
	return nil
}

func (r RelationService) ListFollowInfo(userId uint64) ([]model.User, error) {
	list, err := r.relationManager.ListFollowUser(userId)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (r RelationService) ListFollowerInfo(userId uint64) ([]model.User, error) {
	list, err := r.relationManager.ListFollowerInfo(userId)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (r RelationService) ListFriend(userId uint64) ([]model.User, error) {
	list, err := r.relationManager.ListFriend(userId)
	if err != nil {
		return nil, err
	}
	return list, nil
}
