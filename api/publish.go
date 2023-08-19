package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"tiktok/pkg/util"
)

func Publish(c *gin.Context) {

	file, err := c.FormFile("data")
	if err != nil {
		c.Error(errors.New("文件接收错误"))
		c.Abort()
		return
	}
	util.OSSUpload(file)

}
