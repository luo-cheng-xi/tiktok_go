package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
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
	authorId, err := v.jwtUtil.GetUserIdFromJwt(c.PostForm("token"))
	if err != nil {
		model.AbortWithStatusErrJSON(c, err)
	}
	title := c.PostForm("title")

	file, err := c.FormFile("data")
	if err != nil {
		model.AbortWithStatusErrJSON(c, err)
	}

	//调用service层代码
	err = v.videoService.Publish(file, authorId, title)
	if err != nil {
		model.AbortWithStatusErrJSON(c, err)
	}

	//上传没有出现异常，返回信息
	c.JSON(http.StatusOK, model.NewSuccessRsp())
}
