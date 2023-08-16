package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"tiktok/conf/jwt_conf"
	"tiktok/dao/user_dao"
	"tiktok/model/rsp"
	"tiktok/utils/jwt_utils"
)

// LoginCheckHandler 中间件，用于检查用户是否已经登录
// 中间件必须是gin.HandlerFunc类型
// 包含请求上下文参数
func LoginCheckHandler(c *gin.Context) {
	//截取token字符串的数据载荷
	payload, err := jwt_utils.ParseJwt(c.Query("token"), jwt_conf.JwtSignedKey)
	if err != nil {
		log.Default().Println("错误的jwt令牌 :" + c.Query("token"))
		c.JSON(http.StatusUnauthorized, rsp.Error(rsp.WithMsg("您的用户鉴权无效")))
		c.Abort() //阻止后续中间件的调用
		return    //终止该函数
	}
	//从载荷中提取用户ID信息
	userId := payload.ID
	//查找是否存在该用户信息
	if _, err = user_dao.GetUserById(userId); err != nil {
		c.JSON(http.StatusUnauthorized, rsp.Error(rsp.WithMsg("您的用户鉴权无效")))
		c.Abort() //阻止后续中间件调用
		return    //终止该函数
	}
	//放行
	c.Next()
}
