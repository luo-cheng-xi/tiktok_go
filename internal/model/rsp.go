package model

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/internal/terrs"
)

type BaseRsp struct {
	StatusCode    terrs.StatusCode `json:"status_code"`
	StatusMessage string           `json:"status_message"`
}

type IdAndTokenRsp struct {
	BaseRsp

	Id    uint64 `json:"user_id"`
	Token string `json:"token"`
}

type UserRsp struct {
	BaseRsp

	User UserVO
}

type FeedRsp struct {
	BaseRsp

	NextTime  uint64    `json:"next_time"`
	VideoList []VideoVO `json:"video_list"`
}

type VideoListRsp struct {
	BaseRsp

	VideoList []VideoVO
}

// Option 用于支持Error函数的options模式
type Option func(msg *BaseRsp)

// AbortWithStatusErrJSON 自定义的用于告知前端错误信息的方法
// 传入*gin.Context和error对象，对于自定义类型的错误，会根据
// terrs包中定义的映射推断其http响应码，对于其他非自定义类型的错
// 错误，则统一按照服务器内部异常处理,返回500响应码
func AbortWithStatusErrJSON(c *gin.Context, e error) {
	c.AbortWithStatusJSON(GetHttpCode(e), NewErrorRsp(e))
}
func GetHttpCode(e error) int {
	if terr, ok := e.(terrs.TError); ok {
		//是TError类型，直接包装TError信息
		return terr.GetHttpCode()
	}
	//不是TError类型，返回
	return http.StatusInternalServerError
}

// NewErrorRsp 通过TError信息包装BaseRsp
func NewErrorRsp(e error) BaseRsp {
	if terr, ok := e.(terrs.TError); ok {
		//是TError类型，直接包装TError信息
		return BaseRsp{
			StatusCode:    terr.Code,
			StatusMessage: terr.Error(),
		}
	}
	//不是TError类型，返回
	return BaseRsp{
		StatusCode:    terrs.INTERNAL,
		StatusMessage: e.Error(),
	}
}

// NewSuccessRsp 用于产生表示成功的BaseRsp
func NewSuccessRsp(opt ...Option) BaseRsp {
	resp := BaseRsp{StatusCode: terrs.OK}
	for _, o := range opt {
		o(&resp)
	}
	return resp
}

// WithMsg 用于注入信息
func WithMsg(msg string) Option {
	return func(resp *BaseRsp) {
		resp.StatusMessage = msg
	}
}
