package cmd

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
	"tiktok/internal/model"
	"tiktok/internal/terrs"
)

func InitRouter(r *gin.Engine) {
	con, err := BuildInjector()
	if err != nil {
		debug.PrintStack()
		return
	}
	apiRouter := r.Group("/douyin")
	apiRouter.Use(func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				debug.PrintStack()
				c.AbortWithStatusJSON(http.StatusInternalServerError, model.NewErrorRsp(terrs.ErrInternal))
			}
		}()
		c.Next()
	})
	{
		apiRouter.GET("/user/", con.LoginCheckMiddleware.LoginCheck, con.UserController.GetById)
		apiRouter.POST("/user/register/", con.UserController.Register)
		apiRouter.POST("/user/login/", con.UserController.Login)

		// 根据接口文档，用户投稿token被放在了请求体中，由于投稿实现的第一步就是从token中获取用户id，可以起到登录校验作用，故将只能从请求头中的
		// 登录校验中间件省去
		apiRouter.POST("/publish/action/", con.VideoController.Publish)
	}
}
