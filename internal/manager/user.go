package manager

import (
	"go.uber.org/zap"
	"tiktok/internal/data"
	"tiktok/internal/model"
)

type UserManager struct {
	logger  *zap.Logger
	userDao *data.UserDao
}

func NewUserManager(l *zap.Logger, ud *data.UserDao) *UserManager {
	return &UserManager{
		logger:  l,
		userDao: ud,
	}
}

// GetUserById 通过Id获得用户信息
func (u UserManager) GetUserById(id uint64) (model.User, error) {
	//调用dao层获取用户信息
	user, err := u.userDao.GetUserById(id)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
