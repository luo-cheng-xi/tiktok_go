package user_controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok/model/rsp"
	"tiktok/model/vo"
	"tiktok/service/user_service"
)

// GetById 用户信息获取功能
func GetById(c *gin.Context) {
	//解析参数
	userIdStr := c.Query("user_id")
	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, rsp.Error(rsp.WithMsg(err.Error())))
		return
	}

	//调用service层代码
	userInfo, err := user_service.GetById(uint(userId))
	if err != nil {
		c.JSON(http.StatusOK, rsp.Error(rsp.WithMsg(err.Error())))
		return
	}

	//封装返回值,并返回结果
	response := rsp.UserRsp{
		BaseRsp: rsp.BaseRsp{StatusCode: 0, StatusMessage: "ok"},
		User:    vo.ParseUserVO(userInfo),
	}
	c.JSON(http.StatusOK, response)
}

// Register 注册功能
func Register(c *gin.Context) {
	//解析参数
	username := c.Query("username")
	password := c.Query("password")
	//检查参数是否合法
	if len(username) > 32 {
		c.JSON(http.StatusOK, rsp.Error(rsp.WithMsg("用户名过长")))
	}
	if len(password) <= 5 {
		c.JSON(http.StatusOK, rsp.Error(rsp.WithMsg("用户注册密码过短")))
	}
	if len(password) > 32 {
		c.JSON(http.StatusOK, rsp.Error(rsp.WithMsg("用户注册密码过长")))
	}

	//调用service层代码
	id, token, err := user_service.Register(username, password)
	//该用户已存在，或者出现其他错误
	if err != nil {
		//返回错误信息
		c.JSON(http.StatusOK, rsp.Error(rsp.WithMsg(err.Error())))
		return
	}
	//用户不存在，注册完成，返回id和token
	c.JSON(http.StatusOK, rsp.IdAndTokenRsp{
		BaseRsp: rsp.Success(rsp.WithMsg("注册成功")),
		Id:      id,
		Token:   token,
	})
}

// Login 登录功能
func Login(c *gin.Context) {
	//解析参数
	username := c.Query("username")
	password := c.Query("password")

	//检查参数是否合法
	if len(username) > 32 {
		c.JSON(http.StatusOK, rsp.Error(rsp.WithMsg("用户名过长")))
	}
	if len(password) <= 5 {
		c.JSON(http.StatusOK, rsp.Error(rsp.WithMsg("用户登陆密码过短")))
	}
	if len(password) > 32 {
		c.JSON(http.StatusOK, rsp.Error(rsp.WithMsg("用户登录密码过长")))
	}

	//调用service层代码
	id, token, err := user_service.Login(username, password)
	if err != nil {
		fmt.Print(err.Error())
		c.JSON(http.StatusOK, rsp.Error(rsp.WithMsg(err.Error())))
		return
	}
	//登录信息无误，返回id和token
	c.JSON(http.StatusOK, rsp.IdAndTokenRsp{
		BaseRsp: rsp.Success(rsp.WithMsg("登录成功")),
		Id:      id,
		Token:   token,
	})
}
