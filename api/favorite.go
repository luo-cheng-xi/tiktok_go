package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"tiktok/internal/model"
	"tiktok/internal/service"
	"tiktok/pkg/util"
)

type FavoriteController struct {
	logger             *zap.Logger
	jwtUtil            *util.JwtUtil
	voUtil             *util.VoUtil
	interactionService *service.FavoriteService
}

func NewInteractionController(
	l *zap.Logger,
	ju *util.JwtUtil,
	vu *util.VoUtil,
	is *service.FavoriteService) *FavoriteController {
	return &FavoriteController{
		logger:             l,
		jwtUtil:            ju,
		voUtil:             vu,
		interactionService: is,
	}
}

// FavoriteAction 点赞相关操作
func (v FavoriteController) FavoriteAction(c *gin.Context) {
	//解析参数
	userId, err := v.jwtUtil.GetUserIdFromJwt(c.Query("token"))
	if err != nil {
		v.logger.Debug("无法从token中解析出用户id", zap.String("cause", err.Error()))
		model.AbortWithStatusErrJSON(c, err)
		return
	}
	videoId, err := strconv.ParseUint(c.Query("video_id"), 10, 64)
	if err != nil {
		v.logger.Debug("视频id解析出错", zap.String("cause", err.Error()))
		model.AbortWithStatusErrJSON(c, err)
		return
	}
	actionType, err := strconv.ParseUint(c.Query("action_type"), 10, 32)
	if err != nil {
		v.logger.Debug("actionType解析出错", zap.String("cause", err.Error()))
		model.AbortWithStatusErrJSON(c, err)
		return
	}

	//调用service层代码
	v.interactionService.FavoriteAction(userId, videoId, uint32(actionType))

	//成功，返回信息
	c.JSON(http.StatusOK, model.NewSuccessRsp())
}

// ListFavoriteByUserId 列出用户所有的点赞视频
func (v FavoriteController) ListFavoriteByUserId(c *gin.Context) {
	// 解析参数
	curUserId, err := v.jwtUtil.GetUserIdFromJwt(c.Query("token"))
	if err != nil {
		v.logger.Debug("无法从token中解析用户id", zap.String("cause", err.Error()))
		model.AbortWithStatusErrJSON(c, err)
		return
	}
	tarUserId, err := strconv.ParseUint(c.Query("user_id"), 10, 64)
	if err != nil {
		v.logger.Debug("目标用户id Query参数解析出错", zap.String("cause", err.Error()))
		model.AbortWithStatusErrJSON(c, err)
		return
	}

	//调用service层代码，列出目标用户的所有点赞视频
	videos := v.interactionService.ListFavoriteVideoByUserId(tarUserId)

	//包装为vo类型
	videoVOs := make([]model.VideoVO, len(videos))
	for i, video := range videos {
		videoVOs[i], err = v.voUtil.ParseVideoVO(video, curUserId)
		if err != nil {
			v.logger.Debug("无法从token中解析用户id", zap.String("cause", err.Error()))
			model.AbortWithStatusErrJSON(c, err)
			return
		}
	}

	c.JSON(http.StatusOK, model.VideoListRsp{
		BaseRsp:   model.NewSuccessRsp(),
		VideoList: videoVOs,
	})
}
