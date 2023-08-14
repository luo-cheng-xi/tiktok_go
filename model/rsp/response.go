package rsp

import (
	"tiktok/constant/status_code"
	"tiktok/model/vo"
)

type BaseRsp struct {
	StatusCode    status_code.StatusCode `json:"status_code"`
	StatusMessage string                 `json:"status_message"`
}

type IdAndTokenRsp struct {
	BaseRsp
	Id    uint   `json:"user_id"`
	Token string `json:"token"`
}

type UserRsp struct {
	BaseRsp
	User vo.UserVO
}

// Option 用于支持Error函数的options模式
type Option func(msg *BaseRsp)

// Error 用于产生用于错误的BaseRsp
func Error(opt ...Option) BaseRsp {
	resp := BaseRsp{StatusCode: status_code.Fail}
	for _, o := range opt {
		o(&resp)
	}
	return resp
}

// Success 用于产生表示成功的BaseRsp
func Success(opt ...Option) BaseRsp {
	resp := BaseRsp{StatusCode: status_code.Success}
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
