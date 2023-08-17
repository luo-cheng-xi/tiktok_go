package video_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/lgr"
	"tiktok/model/rsp"
	"tiktok/utils/oss_utils"
)

func Publish(c *gin.Context) {

	file, err := c.FormFile("data")
	if err != nil {
		lgr.Err("文件接收错误")
		c.AbortWithStatusJSON(http.StatusBadRequest, rsp.Error(rsp.WithMsg("文件接收错误")))
		return
	}
	src, err := file.Open()
	if err != nil {
		lgr.Err("文件开启出错")
		c.AbortWithStatusJSON(http.StatusBadRequest, rsp.Error(rsp.WithMsg("无法读取文件")))
		return
	}
	err = oss_utils.Bucket.PutObject("test/testfile"+file.Filename, src)
	if err != nil {
		lgr.Err("文件存储出错")
		c.AbortWithStatusJSON(http.StatusInsufficientStorage, rsp.Error(rsp.WithMsg("文件存储出错")))
		return
	}
	c.JSON(http.StatusOK, rsp.Success(rsp.WithMsg("文件存储成功")))
}
