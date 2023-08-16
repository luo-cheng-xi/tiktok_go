package controller

import "github.com/gin-gonic/gin"

func Publish(c *gin.Context) {
	c.GetPostForm("token")
}
