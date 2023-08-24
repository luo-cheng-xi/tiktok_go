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
	voUtil       *util.VoUtil
	videoService *service.VideoService
}

func NewVideoController(
	zl *zap.Logger,
	uo *util.OssUtil,
	uj *util.JwtUtil,
	vu *util.VoUtil,
	vs *service.VideoService) *VideoController {
	return &VideoController{
		logger:       zl,
		ossUtil:      uo,
		jwtUtil:      uj,
		voUtil:       vu,
		videoService: vs,
	}
}

// Publish 完成视频上传
func (v VideoController) Publish(c *gin.Context) {
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
func (v VideoController) Feed(c *gin.Context) {
	//解析参数
	curUserId, err := v.jwtUtil.GetUserIdFromJwt(c.Query("token"))
	if err != nil {
		v.logger.Debug("用户id信息解析失败", zap.String("cause", err.Error()))
		model.AbortWithStatusErrJSON(c, err)
		return
	}
	latestTimeStamp, err := strconv.ParseInt(c.Query("latest_time"), 10, 64)
	if err != nil {
		v.logger.Debug("字符串转时间错误 :", zap.String("cause", err.Error()))
		model.AbortWithStatusErrJSON(c, err)
		return
	}
	zone := time.FixedZone("CST", 8*3600) //设置为东8区
	latestTime := time.UnixMilli(latestTimeStamp).In(zone)

	//调用service层代码，获取视频信息
	videos, nextTime := v.videoService.Feed(latestTime)

	//封装视频信息
	videoVOs := make([]model.VideoVO, len(videos))
	for i, video := range videos {
		videoVOs[i], err = v.voUtil.ParseVideoVO(video, curUserId)
		if err != nil {
			v.logger.Debug("视频信息转化出错", zap.String("cause", err.Error()))
			model.AbortWithStatusErrJSON(c, err)
			return
		}
	}

	//返回信息
	c.JSON(http.StatusOK, model.FeedRsp{
		BaseRsp:   model.NewSuccessRsp(),
		NextTime:  uint64(nextTime.UnixMilli()),
		VideoList: videoVOs,
	})
}

// ListVideoByAuthorId 列出用户所有投稿过的视频
func (v VideoController) ListVideoByAuthorId(c *gin.Context) {
	//解析参数
	tokenString := c.Query("token")
	userId, err := v.jwtUtil.GetUserIdFromJwt(tokenString)
	authorId, err := strconv.ParseUint(c.Query("user_id"), 10, 64)
	if err != nil {
		v.logger.Debug("作者id解析出错", zap.String("cause", err.Error()))
		model.AbortWithStatusErrJSON(c, err)
		return
	}

	//调用service层代码
	videos := v.videoService.ListVideoByAuthorId(authorId)

	//封装信息为VO
	videoVOs := make([]model.VideoVO, len(videos))
	for i, video := range videos {
		videoVOs[i], err = v.voUtil.ParseVideoVO(video, userId)
		if err != nil {
			model.AbortWithStatusErrJSON(c, err)
			return
		}
	}

	//返回信息
	c.JSON(http.StatusOK, model.VideoListRsp{
		BaseRsp:   model.NewSuccessRsp(),
		VideoList: videoVOs,
	})
}
