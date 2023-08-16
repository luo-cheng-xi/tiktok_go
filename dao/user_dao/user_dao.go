package user_dao

import (
	"errors"
	"gorm.io/gorm"
	"tiktok/dao"
	"tiktok/model/po"
)

func GetUserByUsername(username string) (po.User, error) {
	//查找用户名条件相符的用户
	user := po.User{}
	res := dao.DB.Where(po.User{Username: username}).Take(&user)

	//处理没有找到和异常情况
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return po.User{}, errors.New("没有找到该用户")
		}
		return po.User{}, res.Error
	}

	return user, nil
}

func GetUserById(id uint) (po.User, error) {
	//查找id条件相符的用户
	user := po.User{}
	res := dao.DB.Where("id = ?", id).Take(&user)

	//处理 没有找到的情况 和 异常情况
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return po.User{}, errors.New("没有找到该用户")
		}
		return po.User{}, res.Error
	}

	//返回用户信息
	return user, nil
}
