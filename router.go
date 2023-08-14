package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"tiktok/controller"
)

var logger = log.Default()

func initRouter(r *gin.Engine) {
	apiRouter := r.Group("/douyin")
	{
		apiRouter.GET("/user/", controller.GetUserById)
		apiRouter.POST("/user/register/", controller.Register)
		apiRouter.POST("/user/login/", controller.Login)
	}
}
