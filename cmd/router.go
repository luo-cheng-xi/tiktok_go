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

		apiRouter.POST("/publish/action/", con.LoginCheckMiddleware.LoginCheck, con.VideoController.Publish)
	}
}
