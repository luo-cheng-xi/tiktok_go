package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"tiktok/internal/model"
	"tiktok/internal/service"
	"tiktok/pkg/util"
	"time"
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
		return
	}
	title := c.PostForm("title")

	file, err := c.FormFile("data")
	if err != nil {
		model.AbortWithStatusErrJSON(c, err)
		return
	}

	//调用service层代码
	err = v.videoService.Publish(file, authorId, title)
	if err != nil {
		model.AbortWithStatusErrJSON(c, err)
		return
	}

	//上传没有出现异常，返回信息
	c.JSON(http.StatusOK, model.NewSuccessRsp())
}

// Feed 视频流功能接口
func (v *VideoController) Feed(c *gin.Context) {
	latestTimeStamp, err := strconv.ParseInt(c.Query("latest_time"), 10, 64)
	if err != nil {
		v.logger.Debug("字符串转时间错误 :", zap.String("cause", err.Error()))
		model.AbortWithStatusErrJSON(c, err)
		return
	}
	zone := time.FixedZone("CST", 8*3600) //设置为东8区
	latestTime := time.UnixMilli(latestTimeStamp).In(zone)
	videos, nextTime := v.videoService.Feed(latestTime)
	c.JSON(http.StatusOK, model.FeedRsp{
		BaseRsp:   model.NewSuccessRsp(),
		NextTime:  nextTime.UnixMilli(),
		VideoList: videos,
	})
}
