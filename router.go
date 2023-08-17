package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
	"tiktok/controller/user_controller"
	"tiktok/controller/video_controller"
	"tiktok/handler"
	"tiktok/lgr"
	"tiktok/model/rsp"
)

func initRouter(r *gin.Engine) {
	apiRouter := r.Group("/douyin")
	apiRouter.Use(func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				lgr.Err(fmt.Sprint(err))
				debug.PrintStack()
				c.JSON(http.StatusInternalServerError, rsp.Error(rsp.WithMsg(fmt.Sprint(err))))
				c.Abort()
			}
		}()
		c.Next()
	})
	{
		apiRouter.GET("/user/", handler.LoginCheckHandler, user_controller.GetById)
		apiRouter.POST("/user/register/", user_controller.Register)
		apiRouter.POST("/user/login/", user_controller.Login)
		apiRouter.POST("/publish/action/", video_controller.Publish)

	}
}
