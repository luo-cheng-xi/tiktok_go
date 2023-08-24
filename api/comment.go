package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"tiktok/internal/model"
	"tiktok/internal/service"
	"tiktok/internal/terrs"
	"tiktok/pkg/util"
)

type CommentController struct {
	logger         *zap.Logger
	voUtil         *util.VoUtil
	jwtUtil        *util.JwtUtil
	commentService *service.CommentService
}

func NewCommentController(
	l *zap.Logger,
	vu *util.VoUtil,
	ju *util.JwtUtil,
	cs *service.CommentService) *CommentController {
	return &CommentController{
		logger:         l,
		voUtil:         vu,
		jwtUtil:        ju,
		commentService: cs,
	}
}

// CommentAction 评论操作接口
func (v CommentController) CommentAction(c *gin.Context) {
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
	if actionType != 1 && actionType != 2 {
		model.AbortWithStatusErrJSON(c, terrs.ErrParamInvalid)
		return
	}
	var commentText string
	var commentId uint64
	if actionType == 1 {
		commentText = c.Query("comment_text")
	} else {
		commentId, err = strconv.ParseUint(c.Query("comment_id"), 10, 64)
		if err != nil {
			model.AbortWithStatusErrJSON(c, err)
			return
		}
	}

	//调用service层代码
	comment := v.commentService.CommentAction(userId, videoId, uint32(actionType), commentText, commentId)

	if actionType == 1 {
		commentVO, err := v.voUtil.ParseCommentVO(comment, userId)
		if err != nil {
			model.AbortWithStatusErrJSON(c, err)
			return
		}
		//成功，返回信息
		c.JSON(http.StatusOK, model.CommentRsp{
			BaseRsp: model.NewSuccessRsp(),
			Comment: commentVO,
		})
	} else {
		c.JSON(http.StatusOK, model.NewSuccessRsp())
	}
}

// ListCommentByVideoId 评论列表
func (v CommentController) ListCommentByVideoId(c *gin.Context) {
	userId, err := v.jwtUtil.GetUserIdFromJwt(c.Query("token"))
	if err != nil {
		v.logger.Debug("用户id解析错误", zap.String("cause", err.Error()))
		model.AbortWithStatusErrJSON(c, err)
		return
	}
	videoId, err := strconv.ParseUint(c.Query("video_id"), 10, 64)
	if err != nil {
		v.logger.Debug("视频id解析错误", zap.String("cause", err.Error()))
		model.AbortWithStatusErrJSON(c, terrs.ErrParamInvalid)
		return
	}

	// 调用service 层代码
	comments := v.commentService.ListCommentByVideoId(videoId)

	// 包装并返回内容
	commentVOs := make([]model.CommentVO, len(comments))
	for i, comment := range comments {
		commentVO, err := v.voUtil.ParseCommentVO(comment, userId)
		if err != nil {
			model.AbortWithStatusErrJSON(c, err)
			return
		}
		commentVOs[i] = commentVO
	}
	c.JSON(http.StatusOK, model.CommentListRsp{
		BaseRsp:     model.NewSuccessRsp(),
		CommentList: commentVOs,
	})
}
