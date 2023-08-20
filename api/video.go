package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type VideoController struct {
	logger *zap.Logger
}

func NewVideoController(l *zap.Logger) *VideoController {
	return &VideoController{
		logger: l,
	}
}

func (v *VideoController) Publish(c *gin.Context) {
	//接收文件参数
	//file, err := c.FormFile("data")
	//if err != nil {
	//	c.AbortWithStatusJSON(
	//		http.StatusBadRequest,
	//		model.NewErrorRsp(terrs.ErrInternal))
	//	return
	//}
	//url, err := util.OSSUpload(file)
	//if err != nil {
	//	c.AbortWithStatusJSON(
	//		http.StatusBadRequest,
	//		model.NewErrorRsp(terrs.ErrInternal))
	//	return
	//}
	return
}
