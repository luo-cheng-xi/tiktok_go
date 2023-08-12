package dao

import (
	"tiktok/model"
	"tiktok/utils"
)

func GetUserById(id int) model.User {
	user := model.User{}
	DB.Where("id = ?", id).Find(&user)
	return user
}

func Login(username, password string) (model.User, error) {
	encodedPassword := utils.Encode(password)
	user := model.User{}
	DB.Where("name = ? and password = ?", username, encodedPassword).Find(&user)
	return user, nil
}
