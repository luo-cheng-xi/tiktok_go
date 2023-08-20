package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"tiktok/internal/model"
	"tiktok/internal/service"
	"tiktok/pkg/util"
)

type VideoController struct {
	logger       *zap.Logger
	ossUtil      *util.OssUtil
	jwtUtil      *util.JwtUtil
	videoService *service.VideoService
}

func NewVideoController(
	zl *zap.Logger,
	uo *util.OssUtil,
	uj *util.JwtUtil,
	vs *service.VideoService) *VideoController {
	return &VideoController{
		logger:       zl,
		ossUtil:      uo,
		jwtUtil:      uj,
		videoService: vs,
	}
}

// Publish 完成视频上传
func (v *VideoController) Publish(c *gin.Context) {
	//解析参数
	authorId, err := v.jwtUtil.GetUserIdFromJwt(c.Query("token"))
	if err != nil {
		model.AbortWithStatusErrJSON(c, err)
	}
	file, err := c.FormFile("data")
	if err != nil {
		model.AbortWithStatusErrJSON(c, err)
	}

	//调用service层代码
	err = v.videoService.Publish(file, authorId)
	if err != nil {
		model.AbortWithStatusErrJSON(c, err)
	}
}
