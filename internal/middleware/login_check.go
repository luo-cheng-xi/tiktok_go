package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/internal/model"
	"tiktok/internal/service"
	"tiktok/internal/terrs"
	"tiktok/pkg/util"
)

type LoginCheckMiddleware struct {
	userService *service.UserService
	jwtUtil     *util.JwtUtil
}

func NewLoginCheck(us *service.UserService, ju *util.JwtUtil) *LoginCheckMiddleware {
	return &LoginCheckMiddleware{
		userService: us,
		jwtUtil:     ju,
	}
}

// LoginCheck 中间件，用于检查用户是否已经登录
// 中间件必须是gin.HandlerFunc类型
// 包含请求上下文参数
func (rx *LoginCheckMiddleware) LoginCheck(c *gin.Context) {
	//截取token字符串的数据载荷
	payload, err := rx.jwtUtil.ParseJwt(c.Query("token"))
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			model.NewErrorRsp(err))
		return //终止该函数
	}
	//从载荷中提取用户ID信息
	userId := payload.ID
	//查找是否存在该用户信息
	if _, err = rx.userService.GetById(userId); err != nil {
		if terrs.ErrUserNotFound.Eq(err) {
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				model.NewErrorRsp(terrs.ErrTokenInvalid))
			return //终止该函数
		}
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			model.NewErrorRsp(terrs.ErrInternal))
		return
	}
	//放行
	c.Next()
}
