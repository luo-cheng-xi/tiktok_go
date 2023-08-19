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
		c.Error(err)
		c.Abort()
		return
	}
	//调用service层代码
	userInfo, err := service.GetById(uint(userId))
	if err != nil {
		u.logger.Error("service.GetById() error : ", zap.String("detail", err.Error()))
		c.Error(err)
		c.Abort()
		return
	}
	//封装返回值,并返回结果
	c.JSON(http.StatusOK, model.UserRsp{
		BaseRsp: model.NewSuccessRsp(),
		User:    model.ParseUserVO(userInfo),
	})
}

// Register 注册功能
func (*UserController) Register(c *gin.Context) {
	//解析参数
	username := c.Query("username")
	password := c.Query("password")

	//检查参数是否合法
	if len(username) > 32 {
		c.Error(terrs.ErrUsernameTooLong)
		c.Abort()
		return
	}
	if len(password) <= 5 {
		c.Error(terrs.ErrPasswordTooShort)
		c.Abort()
		return
	}
	if len(password) > 32 {
		c.Error(terrs.ErrPasswordTooLong)
		c.Abort()
		return
	}

	//调用service层代码
	id, token, err := service.Register(username, password)
	//该用户已存在，或者出现其他错误
	if err != nil {
		//返回错误信息
		c.Error(err)
		c.Abort()
		return
	}
	//用户不存在，注册完成，返回id和token
	c.JSON(http.StatusOK, model.IdAndTokenRsp{
		BaseRsp: model.NewSuccessRsp(),
		Id:      id,
		Token:   token,
	})
}

// Login 登录功能
func (*UserController) Login(c *gin.Context) {
	//解析参数
	username := c.Query("username")
	password := c.Query("password")

	//检查参数是否合法
	if len(username) > 32 {
		c.Error(terrs.ErrUsernameTooLong)
		c.Abort()
		return
	}
	if len(password) <= 5 {
		c.Error(terrs.ErrPasswordTooShort)
		c.Abort()
		return
	}
	if len(password) > 32 {
		c.Error(terrs.ErrPasswordTooShort)
		c.Abort()
		return
	}

	//调用service层代码
	id, token, err := service.Login(username, password)
	if err != nil {
		//对于没有找到用户和没有找到密码的情况，进行处理。其他异常情况直接作为内部异常告知前端
		if terrs.ErrUserNotFound.Eq(err) || terrs.ErrPasswordWrong.Eq(err) {
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				model.NewErrorRsp(terrs.INTERNAL, model.WithMsg(err.Error())))
		} else {
			c.AbortWithStatusJSON(
				http.StatusInternalServerError,
				model.NewErrorRsp(terrs.INTERNAL, model.WithMsg(terrs.ErrInternal.Error())))
		}
		return
	}

	//登录信息无误，返回id和token
	c.JSON(http.StatusOK, model.IdAndTokenRsp{
		BaseRsp: model.NewSuccessRsp(),
		Id:      id,
		Token:   token,
	})
}
