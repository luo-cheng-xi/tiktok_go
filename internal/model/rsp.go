package model

import (
	"tiktok/internal/terrs"
)

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
