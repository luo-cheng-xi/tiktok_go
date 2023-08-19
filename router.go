package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
	"tiktok/api"
	"tiktok/internal/middleware"
	"tiktok/internal/model"
	"tiktok/internal/terrs"
)

func initRouter(r *gin.Engine) {

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
		userController := api.NewUserController()
		apiRouter.GET("/user/", middleware.LoginCheckHandler, userController.GetById)
		apiRouter.POST("/user/register/", userController.Register)
		apiRouter.POST("/user/login/", userController.Login)

		apiRouter.POST("/publish/action/", api.Publish)
	}
}
