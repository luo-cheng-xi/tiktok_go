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

type RelationController struct {
	logger          *zap.Logger
	relationService *service.RelationService
	jwtUtil         *util.JwtUtil
	voUtil          *util.VoUtil
}

func NewRelationController(
	l *zap.Logger,
	rs *service.RelationService,
	ju *util.JwtUtil,
	vo *util.VoUtil) *RelationController {
	return &RelationController{
		logger:          l,
		relationService: rs,
		jwtUtil:         ju,
		voUtil:          vo,
	}
}

func (r RelationController) FollowAction(c *gin.Context) {
	//解析参数
	userId, err := r.jwtUtil.GetUserIdFromJwt(c.Query("token"))
	if err != nil {
		r.logger.Debug("用户id解析错误", zap.String("cause", err.Error()))
		model.AbortWithStatusErrJSON(c, err)
		return
	}
	toUserId, err := strconv.ParseUint(c.Query("to_user_id"), 10, 64)
	if err != nil {
		r.logger.Debug("参数解析错误", zap.String("cause", err.Error()))
		model.AbortWithStatusErrJSON(c, err)
		return
	}
	actionType, err := strconv.ParseUint(c.Query("action_type"), 10, 64)
	if err != nil {
		r.logger.Debug("参数解析错误", zap.String("cause", err.Error()))
		model.AbortWithStatusErrJSON(c, err)
		return
	}
	// actionType 为1关注，为2取关，都不是说明有问题，需要处理
	if actionType != 1 && actionType != 2 {
		model.AbortWithStatusErrJSON(c, terrs.ErrParamInvalid)
		return
	}

	//调用service层代码
	err = r.relationService.FollowAction(userId, toUserId, actionType)
	if err != nil {
		model.AbortWithStatusErrJSON(c, err)
		return
	}

	c.JSON(http.StatusOK, model.NewSuccessRsp())
}

func (r RelationController) ListFollow(c *gin.Context) {
	// 解析参数
	curUserId, err := r.jwtUtil.GetUserIdFromJwt(c.Query("token"))
	if err != nil {
		r.logger.Debug("解析当前用户id失败", zap.String("cause", err.Error()))
		model.AbortWithStatusErrJSON(c, err)
		return
	}
	tarUserId, err := strconv.ParseUint(c.Query("user_id"), 10, 64)
	if err != nil {
		r.logger.Debug("目标用户id解析错误", zap.String("cause", err.Error()))
		model.AbortWithStatusErrJSON(c, err)
		return
	}

	//调用service层代码，获取该用户关注的所有用户信息
	users, err := r.relationService.ListFollowInfo(tarUserId)
	if err != nil {
		model.AbortWithStatusErrJSON(c, err)
		return
	}

	//封装并返回结果
	var userVOs = make([]model.UserVO, len(users))
	for i, user := range users {
		userVOs[i], err = r.voUtil.ParseUserVO(user, curUserId)
		if err != nil {
			model.AbortWithStatusErrJSON(c, err)
			return
		}
	}

	c.JSON(http.StatusOK, model.UserListRsp{
		BaseRsp:  model.NewSuccessRsp(),
		UserList: userVOs,
	})
}

func (r RelationController) ListFollower(c *gin.Context) {
	//解析参数
	curUserId, err := r.jwtUtil.GetUserIdFromJwt(c.Query("token"))
	if err != nil {
		r.logger.Debug("解析当前用户id失败", zap.String("cause", err.Error()))
		model.AbortWithStatusErrJSON(c, err)
		return
	}
	tarUserId, err := strconv.ParseUint(c.Query("user_id"), 10, 64)
	if err != nil {
		r.logger.Debug("目标用户id解析错误", zap.String("cause", err.Error()))
		model.AbortWithStatusErrJSON(c, err)
		return
	}

	//调用service层代码，获取该用户关注的所有粉丝信息
	users, err := r.relationService.ListFollowerInfo(tarUserId)
	if err != nil {
		model.AbortWithStatusErrJSON(c, err)
		return
	}

	var userVOs = make([]model.UserVO, len(users))
	for i, user := range users {
		userVOs[i], err = r.voUtil.ParseUserVO(user, curUserId)
		if err != nil {
			model.AbortWithStatusErrJSON(c, err)
			return
		}
	}

	c.JSON(http.StatusOK, model.UserListRsp{
		BaseRsp:  model.NewSuccessRsp(),
		UserList: userVOs,
	})
}

func (r RelationController) ListFriend(c *gin.Context) {
	// 解析参数
	curUserId, err := r.jwtUtil.GetUserIdFromJwt(c.Query("token"))
	if err != nil {
		r.logger.Debug("解析当前用户id失败", zap.String("cause", err.Error()))
		model.AbortWithStatusErrJSON(c, err)
		return
	}
	//解析参数
	userId, err := strconv.ParseUint(c.Query("user_id"), 10, 64)
	if err != nil {
		r.logger.Debug("参数解析失败", zap.String("cause", err.Error()))
		model.AbortWithStatusErrJSON(c, err)
		return
	}

	//调用service层代码获取好友信息
	users, err := r.relationService.ListFriend(userId)
	if err != nil {
		model.AbortWithStatusErrJSON(c, err)
		return
	}

	userVOs := make([]model.UserVO, len(users))
	for i, user := range users {
		userVOs[i], err = r.voUtil.ParseUserVO(user, curUserId)
		if err != nil {
			r.logger.Debug("转化VO出错", zap.String("cause", err.Error()))
			model.AbortWithStatusErrJSON(c, err)
			return
		}
	}

	c.JSON(http.StatusOK, model.UserListRsp{
		BaseRsp:  model.NewSuccessRsp(),
		UserList: userVOs,
	})
}
