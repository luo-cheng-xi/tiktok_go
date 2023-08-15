package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"tiktok/conf"
	"tiktok/dao"
	"tiktok/utils"
)

// LoginCheckHandler 中间件，用于检查用户是否已经登录
// 中间件必须是gin.HandlerFunc类型
// 包含请求上下文参数
func LoginCheckHandler(c *gin.Context) {
	//截取token字符串的数据载荷
	payload, err := utils.ParseJwt(c.Query("token"), conf.JwtSignedKey)
	if err != nil {
		log.Default().Println("错误的jwt令牌 :" + c.Query("token"))
		return
	}
	//从载荷中提取用户名信息
	username := payload.Username
	//查找是否存在该用户信息
	if _, err = dao.GetUserByUsername(username); err != nil {
		log.Default().Println("jwt令牌负载的用户信息不存在 用户名为 : " + username)
		return
	}
	//放行
	c.Next()
}
