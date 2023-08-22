package service

import (
	"go.uber.org/zap"
	"tiktok/internal/data"
	"tiktok/internal/manager"
	"tiktok/internal/model"
	"tiktok/internal/terrs"
	"tiktok/pkg/util"
)

type UserService struct {
	logger      *zap.Logger
	userDao     *data.UserDao
	videoDao    *data.VideoDao
	favoriteDao *data.FavoriteDao
	relationDao *data.RelationDao
	manager     *manager.Manager
	jwtUtil     *util.JwtUtil
}

func NewUserService(
	l *zap.Logger,
	ud *data.UserDao,
	rd *data.RelationDao,
	vd *data.VideoDao,
	fd *data.FavoriteDao,
	m *manager.Manager,
	ju *util.JwtUtil) *UserService {
	return &UserService{
		logger:      l,
		userDao:     ud,
		relationDao: rd,
		videoDao:    vd,
		favoriteDao: fd,
		manager:     m,
		jwtUtil:     ju,
	}
}

// Register 用户注册功能
//
// error : ErrUsernameRegistered
func (u UserService) Register(username, password string) (uint64, string, error) {
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
func (u UserService) Login(username, password string) (uint64, string, error) {
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
func (u UserService) GetUserById(id uint64) (model.User, error) {
	//调用dao层获取用户信息
	user, err := u.userDao.GetUserById(id)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

// GetFollowCount 通过用户Id获得用户关注了的用户的人数
func (u UserService) GetFollowCount(userId uint64) uint64 {
	return u.relationDao.GetFollowCount(userId)
}

// GetFollowerCount 通过用户Id获得用户的粉丝数
func (u UserService) GetFollowerCount(userId uint64) uint64 {
	return u.relationDao.GetFollowerCount(userId)
}

// IsFollow 通过关注者和被关注者的id获取是否存在该关注关系
func (u UserService) IsFollow(userId uint64, toUserId uint64) (bool, error) {
	flag, err := u.relationDao.IsFollow(userId, toUserId)
	//出错则返回是否存在关注关系
	if err != nil {
		return false, err
	}
	return flag, nil
}

// ParseUserVO 待补全 ,将user转化为UserVO
func (u UserService) ParseUserVO(tarUser model.User, curUserId uint64) (model.UserVO, error) {
	isFollow, err := u.IsFollow(curUserId, tarUser.ID)
	if err != nil {
		return model.UserVO{}, err
	}
	return model.UserVO{
		ID:              tarUser.ID,
		Name:            tarUser.Username,
		FollowCount:     u.GetFollowCount(tarUser.ID),
		FollowerCount:   u.GetFollowerCount(tarUser.ID),
		IsFollow:        isFollow,
		Avatar:          tarUser.Avatar,
		BackgroundImage: tarUser.BackgroundImage,
		Signature:       tarUser.Signature,
		TotalFavorited:  u.manager.GetTotalFavorited(tarUser.ID),
		WorkCount:       u.videoDao.CountVideoByAuthorId(tarUser.ID),
		FavoriteCount:   u.favoriteDao.CountFavoriteByUserId(tarUser.ID),
	}, nil
}
