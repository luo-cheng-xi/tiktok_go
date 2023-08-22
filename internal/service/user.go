package service

import (
	"go.uber.org/zap"
	"tiktok/internal/data"
	"tiktok/internal/model"
	"tiktok/internal/terrs"
	"tiktok/pkg/util"
)

type UserService struct {
	logger  *zap.Logger
	userDao *data.UserDao
	jwtUtil *util.JwtUtil
}

func NewUserService(
	l *zap.Logger,
	ud *data.UserDao,
	ju *util.JwtUtil) *UserService {
	return &UserService{
		logger:  l,
		userDao: ud,
		jwtUtil: ju,
	}
}

// Register 用户注册功能
//
// error : ErrUsernameRegistered
func (u *UserService) Register(username, password string) (uint64, string, error) {
	_, err := u.userDao.GetUserByUsername(username)
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
	//返回用户的id和token
	userId := u.userDao.CreateUser(user)
	token := u.jwtUtil.GetJwt(userId)
	return userId, token, nil
}

// Login 用户登录功能
//
// error: ErrPasswordWrong | ErrUserNotFound
func (u *UserService) Login(username, password string) (uint64, string, error) {
	//查找是否存在该用户名的用户
	user, err := u.userDao.GetUserByUsername(username)
	if err != nil {
		return 0, "", err
	}

	//校验密码是否正确
	flag, err := util.MatchPasswordAndHash(password, user.Password)
	//如果出错，抛错误；如果不正确，抛ErrPasswordWrong
	if err != nil {
		return 0, "", err
	} else if !flag {
		return 0, "", terrs.ErrPasswordWrong
	}
	//密码正确，返回用户id,token令牌，nil
	return user.ID, u.jwtUtil.GetJwt(user.ID), nil
}

// GetUserById 通过Id获得用户信息
//
// error : ErrUserNotFound
func (u *UserService) GetUserById(id uint64) (model.User, error) {
	//调用dao层获取用户信息
	user, err := u.userDao.GetUserById(id)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

// GetFollowCount 通过用户Id获得用户关注人数
func GetFollowCount(userId uint64) {
	data.GetFollowCount(userId)
}
