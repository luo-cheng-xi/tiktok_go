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
				c.AbortWithStatusJSON(http.StatusInternalServerError, model.NewErrorRsp(terrs.ErrInternal))
			}
		}()
		c.Next()
	})
	{
		// 用户信息
		apiRouter.GET("/user/", con.LoginCheckMiddleware.LoginCheck, con.UserController.GetUserById)
		// 用户注册
		apiRouter.POST("/user/register/", con.UserController.Register)
		// 用户登录
		apiRouter.POST("/user/login/", con.UserController.Login)

		// 投稿接口，根据接口文档，用户投稿token被放在了请求体中，由于投稿实现的第一步就是从token中获取用户id，可以起到登录校验作用，故将只能从请求头中的登录校验中间件省去
		apiRouter.POST("/publish/action/", con.VideoController.Publish)
		// 视频流接口
		apiRouter.GET("/feed/", con.LoginCheckMiddleware.LoginCheck, con.VideoController.Feed)

		// 发布列表
		apiRouter.GET("/publish/list/", con.LoginCheckMiddleware.LoginCheck, con.VideoController.ListVideoByAuthorId)

		// 赞操作
		apiRouter.POST("/favorite/action/", con.LoginCheckMiddleware.LoginCheck, con.favoriteController.FavoriteAction)

		// 喜欢列表
		apiRouter.GET("/favorite/list/", con.LoginCheckMiddleware.LoginCheck, con.favoriteController.ListFavoriteByUserId)

		// 关注操作
		apiRouter.POST("/relation/action/", con.LoginCheckMiddleware.LoginCheck, con.relationController.FollowAction)

		// 关注列表
		apiRouter.GET("/relation/follow/list/", con.LoginCheckMiddleware.LoginCheck, con.relationController.ListFollow)

		// 好友列表
		apiRouter.GET("/relation/friend/list/", con.LoginCheckMiddleware.LoginCheck, con.relationController.ListFriend)

		// 粉丝列表
		apiRouter.GET("/relation/follower/list/", con.LoginCheckMiddleware.LoginCheck, con.relationController.ListFollower)
	}
}
