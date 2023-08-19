package dao

import (
	"errors"
	"gorm.io/gorm"
	"tiktok/internal/model"
	"tiktok/internal/terrs"
)

/*
GetUserByUsername 通过用户名获取用户信息
Param:

	username 用户名

Return:

	po.User 用户信息结构体
	error   错误
*/
func GetUserByUsername(username string) (model.User, error) {
	//查找用户名条件相符的用户
	user := model.User{}
	err := DB.Where("username = ?", username).Take(&user).Error

	//异常处理
	if err != nil {
		//没有找到该用户
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.User{}, terrs.ErrUserNotFound
		}
		//出现未知的异常
		return model.User{}, terrs.ErrInternal
	}

	return user, nil
}

func GetUserById(id uint) (model.User, error) {
	//查找id条件相符的用户
	user := model.User{}
	err := DB.Where("id = ?", id).Take(&user).Error

	//处理 没有找到的情况 和 异常情况
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.User{}, terrs.ErrUserNotFound
		}
		//出现未知的异常
		return model.User{}, terrs.ErrInternal
	}

	//返回用户信息
	return user, nil
}
