package main

import (
	"tiktok/dao"
	"tiktok/model"
)

func GetUserById(id int) model.User {
	user := dao.GetUserById(id)
	return user
}

// Login returns true while username exists and password is correct
// otherwise it will throw an error
func Login(username, password string) (model.User, error) {
	user, err := dao.Login(username, password)
	if err != nil {
		return user, err
	}
	return user, nil
}
