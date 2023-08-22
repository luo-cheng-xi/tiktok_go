package manager

import (
	"go.uber.org/zap"
	"tiktok/internal/data"
)

type VideoManager struct {
	logger   *zap.Logger
	videoDao *data.VideoDao
}

func NewVideoManager(zl *zap.Logger, vd *data.VideoDao) *VideoManager {
	return &VideoManager{
		logger:   zl,
		videoDao: vd,
	}
}
