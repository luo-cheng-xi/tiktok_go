package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"tiktok/controller/user_controller"
	"tiktok/controller/video_controller"
	"tiktok/handler"
)

var logger = log.Default()

func initRouter(r *gin.Engine) {
	apiRouter := r.Group("/douyin")
	{
		apiRouter.GET("/user/", handler.LoginCheckHandler, user_controller.GetById)
		apiRouter.POST("/user/register/", user_controller.Register)
		apiRouter.POST("/user/login/", user_controller.Login)
		apiRouter.POST("/public/action/", video_controller.Publish)
	}
}
