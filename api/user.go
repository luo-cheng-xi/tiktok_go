package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"tiktok/internal/model"
	"tiktok/internal/service"
	"tiktok/internal/terrs"
)

type UserController struct {
	logger      *zap.Logger
	userService *service.UserService
}

func NewUserController(l *zap.Logger, us *service.UserService) *UserController {
	return &UserController{
		logger:      l,
		userService: us,
	}
}

// GetById 用户信息获取功能
func (rx *UserController) GetById(c *gin.Context) {
	//解析参数
	userIdStr := c.Query("user_id")
	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		rx.logger.Error("strconv.ParseUint error : ", zap.String("detail", err.Error()))
		model.AbortWithStatusErrJSON(c, err)
		return
	}

	//调用service层代码
	userInfo, err := rx.userService.GetById(uint(userId))
	if err != nil {
		model.AbortWithStatusErrJSON(c, err)
		return
	}

	//封装返回值,并返回结果
	c.JSON(http.StatusOK, model.UserRsp{
		BaseRsp: model.NewSuccessRsp(),
		User:    model.ParseUserVO(userInfo),
	})
}

// Register 注册功能
func (rx *UserController) Register(c *gin.Context) {
	//解析参数
	username := c.Query("username")
	password := c.Query("password")

	//检查参数是否合法
	if len(username) > 32 {
		model.AbortWithStatusErrJSON(c, terrs.ErrUsernameTooLong)
		return
	}
	if len(password) <= 5 {
		model.AbortWithStatusErrJSON(c, terrs.ErrPasswordTooShort)
		return
	}
	if len(password) > 32 {
		model.AbortWithStatusErrJSON(c, terrs.ErrPasswordTooLong)
		return
	}

	//调用service层代码
	id, token, err := rx.userService.Register(username, password)
	if err != nil {
		model.AbortWithStatusErrJSON(c, err)
		return
	}

	//用户不存在，注册完成，返回id和token
	c.JSON(
		http.StatusOK,
		model.IdAndTokenRsp{
			BaseRsp: model.NewSuccessRsp(),
			Id:      id,
			Token:   token,
		})
}

// Login 登录功能
func (rx *UserController) Login(c *gin.Context) {
	//解析参数
	username := c.Query("username")
	password := c.Query("password")

	//检查参数是否合法
	if len(username) > 32 {
		model.AbortWithStatusErrJSON(c, terrs.ErrUsernameTooLong)
		return
	}
	if len(password) <= 5 {
		model.AbortWithStatusErrJSON(c, terrs.ErrPasswordTooShort)
		return
	}
	if len(password) > 32 {
		model.AbortWithStatusErrJSON(c, terrs.ErrPasswordTooLong)
		return
	}

	//调用service层代码
	id, token, err := rx.userService.Login(username, password)
	if err != nil {
		model.AbortWithStatusErrJSON(c, err)
		return
	}

	//登录信息无误，返回id和token
	c.JSON(
		http.StatusOK,
		model.IdAndTokenRsp{
			BaseRsp: model.NewSuccessRsp(),
			Id:      id,
			Token:   token,
		})
}
