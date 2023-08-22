package service

import "go.uber.org/zap"

type RelationService struct {
	logger *zap.Logger
}

func NewRelationService(zl *zap.Logger) *RelationService {
	return &RelationService{
		logger: zl,
	}
}
