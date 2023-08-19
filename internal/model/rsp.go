package model

import "tiktok/internal/terrs"

type BaseRsp struct {
	StatusCode    terrs.StatusCode `json:"status_code"`
	StatusMessage string           `json:"status_message"`
}

type IdAndTokenRsp struct {
	BaseRsp
	Id    uint   `json:"user_id"`
	Token string `json:"token"`
}

type UserRsp struct {
	BaseRsp
	User UserVO
}

// Option 用于支持Error函数的options模式
type Option func(msg *BaseRsp)

// NewErrorRsp 用于产生用于错误的BaseRsp
func NewErrorRsp(code terrs.StatusCode, opt ...Option) BaseRsp {
	resp := BaseRsp{StatusCode: code}
	for _, o := range opt {
		o(&resp)
	}
	return resp
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
