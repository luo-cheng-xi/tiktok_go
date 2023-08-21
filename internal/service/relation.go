package service

import "go.uber.org/zap"

type RelationService struct {
	logger *zap.Logger
}

func NewRelationController(zl *zap.Logger) *RelationService {
	return &RelationService{
		logger: zl,
	}
}
