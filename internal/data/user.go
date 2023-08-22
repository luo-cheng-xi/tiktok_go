package data

import (
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"tiktok/internal/model"
	"tiktok/internal/terrs"
)

type UserDao struct {
	logger *zap.Logger
	db     *gorm.DB
}

func NewUserDao(l *zap.Logger, d *Data) *UserDao {
	return &UserDao{
		logger: l,
		db:     d.DB,
	}
}

// GetUserByUsername 通过用户名获取用户信息
//
// error : ErrorUserNotFound
func (rx *UserDao) GetUserByUsername(username string) (model.User, error) {
	//查找用户名条件相符的用户
	user := model.User{}
	err := rx.db.Where("username = ?", username).Take(&user).Error
	//异常处理
	if err != nil {
		//没有找到该用户
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.User{}, terrs.ErrUserNotFound
		}
		//出现未知的异常
		return model.User{}, err
	}
	return user, nil
}

// GetUserById 通过用户Id获取用户信息
//
// error : ErrUserNotFound
func (rx *UserDao) GetUserById(id uint64) (model.User, error) {
	//查找id条件相符的用户
	user := model.User{}
	err := rx.db.Where("id = ?", id).Take(&user).Error

	//处理 没有找到的情况 和 异常情况
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.User{}, terrs.ErrUserNotFound
		}
		//出现未知的异常
		return model.User{}, err
	}

	//返回用户信息
	return user, nil
}

// CreateUser 创建用户，并返回主键id
func (rx *UserDao) CreateUser(user model.User) uint64 {
	rx.db.Create(&user)
	return user.ID
}
