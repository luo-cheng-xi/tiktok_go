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
	logger *zap.Logger
}

func NewUserController() *UserController {
	return &UserController{
		logger: zap.NewExample(),
	}
}

// GetById 用户信息获取功能
func (u *UserController) GetById(c *gin.Context) {
	//解析参数
	userIdStr := c.Query("user_id")
	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		u.logger.Error("strconv.ParseUint error : ", zap.String("detail", err.Error()))

		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			model.NewErrorRsp(terrs.ErrInternal))
		return
	}

	//调用service层代码
	userInfo, err := service.GetById(uint(userId))
	if err != nil {
		if terrs.ErrUserNotFound.Eq(err) {
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				model.NewErrorRsp(terrs.ErrUserNotFound))
			return
		}
		u.logger.Error("service.GetById() error : ", zap.String("detail", err.Error()))
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			model.NewErrorRsp(terrs.ErrInternal))
		return
	}

	//封装返回值,并返回结果
	c.JSON(http.StatusOK, model.UserRsp{
		BaseRsp: model.NewSuccessRsp(),
		User:    model.ParseUserVO(userInfo),
	})
}

// Register 注册功能
func (u *UserController) Register(c *gin.Context) {
	//解析参数
	username := c.Query("username")
	password := c.Query("password")

	//检查参数是否合法
	if len(username) > 32 {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			model.NewErrorRsp(terrs.ErrUsernameTooLong))
		return
	}
	if len(password) <= 5 {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			model.NewErrorRsp(terrs.ErrPasswordTooShort))
		return
	}
	if len(password) > 32 {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			model.NewErrorRsp(terrs.ErrPasswordTooLong))
		return
	}

	//调用service层代码
	id, token, err := service.Register(username, password)
	//该用户已存在，或者出现其他错误
	if err != nil {
		if terrs.ErrUsernameRegistered.Eq(err) {
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				model.NewErrorRsp(err))
			return
		}
		u.logger.Error("service.Register error : ", zap.String("detail", err.Error()))
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			model.NewErrorRsp(terrs.ErrInternal))
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
func (u *UserController) Login(c *gin.Context) {
	//解析参数
	username := c.Query("username")
	password := c.Query("password")

	if len(username) > 32 {
		// 用户名过长，返回错误信息
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			model.NewErrorRsp(terrs.ErrUsernameTooLong))
		return
	}
	if len(password) <= 5 {
		// 密码过短，返回错误信息
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			model.NewErrorRsp(terrs.ErrPasswordTooShort))
		return
	}
	if len(password) > 32 {
		// 密码过长，返回错误信息
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			model.NewErrorRsp(terrs.ErrPasswordTooLong))
		return
	}

	//调用service层代码
	id, token, err := service.Login(username, password)
	if err != nil {
		if terrs.ErrUserNotFound.Eq(err) || terrs.ErrPasswordWrong.Eq(err) {
			//对于找不到该用户和密码错误的情况，将错误信息告知前端
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				model.NewErrorRsp(err.(terrs.TError)))
			return
		}
		//对于其它的服务器内部出现的错误，告知前端服务器存在内部错误，在控制台打印日志
		u.logger.Error("service.Login error : ", zap.String("error", err.Error()))
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			model.NewErrorRsp(terrs.ErrInternal))

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
