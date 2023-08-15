package service

import (
	"errors"
	"log"
	"tiktok/dao"
	"tiktok/model/po"
	"tiktok/utils"
)

/*
Register 用于支持controller中的注册功能
参数

	username 用户名
	password 密码

返回值

	uint 用户id
	string 用户权限token
	error 错误
*/
func Register(username, password string) (uint, string, error) {
	_, err := dao.GetUserByUsername(username)
	//err == nil时，说明通过用户名找到了该用户，返回
	if err == nil {
		return 0, "", errors.New("该用户已存在")
	}
	//用户不存在，执行创建逻辑
	//首先对用户密码进行加密
	encodePassword := utils.Encode(password)
	//存储该用户信息
	user := po.User{
		Username: username,
		Password: encodePassword,
	}
	dao.DB.Create(&user)
	//返回用户的id和token
	id := user.ID
	token := utils.GetJwt(user.Username)
	return id, token, nil
}

func Login(username, password string) (uint, string, error) {
	//查找是否存在该用户名的用户
	user, err := dao.GetUserByUsername(username)
	if err != nil {
		return 0, "", err
	}

	//校验密码是否正确
	if !utils.Check(password, user.Password) {
		return 0, "", errors.New("密码错误")
	}
	//密码正确，返回用户id,token令牌，nil
	return user.ID, utils.GetJwt(username), nil
}

// GetUserById 获
func GetUserById(id int64) (po.User, error) {
	//调用dao层获取用户信息
	user, err := dao.GetUserById(id)
	if err != nil {
		log.Default().Println(err)
		return po.User{}, err
	}
	return user, nil
}
