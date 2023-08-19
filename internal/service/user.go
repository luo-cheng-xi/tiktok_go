package service

import (
	"log"
	"tiktok/internal/dao"
	"tiktok/internal/model"
	"tiktok/internal/terrs"
	"tiktok/pkg/util"
)

// Register 用于支持controller中的注册功能
func Register(username, password string) (uint, string, error) {
	_, err := dao.GetUserByUsername(username)
	//err == nil时，说明通过用户名找到了该用户，返回
	if err == nil {
		return 0, "", terrs.ErrUsernameRegistered
	}

	//用户不存在，执行创建逻辑
	//首先对用户密码进行加密
	encodePassword, err := util.EncryptPassword(password)
	if err != nil {
		return 0, "", err
	}
	//存储该用户信息
	user := model.User{
		Username: username,
		Password: encodePassword,
	}
	dao.DB.Create(&user)
	//返回用户的id和token
	id := user.ID
	token := util.GetJwt(id)
	return id, token, nil
}

/*
Login
Param

	username
	password

Return

	uint 用户id
	string token令牌
	error
		可能返回自定义的
		ErrUserNotFound
		ErrPasswordWrong
		或者内部异常
*/
func Login(username, password string) (uint, string, error) {
	//查找是否存在该用户名的用户
	user, err := dao.GetUserByUsername(username)
	if err != nil {
		return 0, "", terrs.ErrUserNotFound
	}

	//校验密码是否正确
	flag, err := util.MatchPasswordAndHash(password, user.Password)
	if !flag || err != nil {
		return 0, "", terrs.ErrPasswordWrong
	}
	//密码正确，返回用户id,token令牌，nil
	return user.ID, util.GetJwt(user.ID), nil
}

// GetById 获
func GetById(id uint) (model.User, error) {
	//调用dao层获取用户信息
	user, err := dao.GetUserById(id)
	if err != nil {
		log.Default().Println(err)
		return model.User{}, err
	}
	return user, nil
}

func GetFollowCount(userId uint) {
	dao.GetFollowCount(userId)
}
