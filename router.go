package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"tiktok/dao"
)

func initRouter(r *gin.Engine) {
	apiRouter := r.Group("/douyin")
	apiRouter.GET("/user/", func(c *gin.Context) {
		userIdStr := c.Query("user_id")
		userId, _ := strconv.Atoi(userIdStr)
		user := dao.GetUserById(userId)
		c.JSON(200, user)
	},
	)
}
