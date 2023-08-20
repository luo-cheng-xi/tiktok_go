package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"tiktok/internal/model"
	"tiktok/internal/service"
	"tiktok/internal/terrs"
	"tiktok/pkg/util"
)

type VideoController struct {
	logger       *zap.Logger
	ossUtil      *util.OssUtil
	videoService *service.VideoService
}

func NewVideoController(l *zap.Logger, o *util.OssUtil) *VideoController {
	return &VideoController{
		logger:  l,
		ossUtil: o,
	}
}

// Publish 完成视频上传
func (v *VideoController) Publish(c *gin.Context) {
	//接收文件参数
	file, err := c.FormFile("data")
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			model.NewErrorRsp(terrs.ErrInternal))
		return
	}
	err = v.videoService.Publish(file)
	if err != nil {

	}
}
