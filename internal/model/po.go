package model

import (
	"fmt"
	"gorm.io/gorm"
	"tiktok/internal/conf"
	"time"
)

var defaultConf = conf.GetDefaultConf()

// User 用户实体类
// 唯一联合索引 Username -> Password
type User struct {
	gorm.Model
	Username        string `gorm:"size:32;uniqueIndex:idx_username_password"`
	Password        string `gorm:"size:120;uniqueIndex:idx_username_password"`
	Avatar          string `gorm:"size:255"`
	BackgroundImage string `gorm:"size:255"`
	Signature       string `gorm:"size:200"`
}

func (u *User) Reg() {
	fmt.Println("Reg")
}

// BeforeCreate Hook函数，用户设置用户的默认信息
func (u *User) BeforeCreate(*gorm.DB) error {
	if u.Avatar == "" {
		u.Avatar = defaultConf.UserAvatar
	}
	if u.BackgroundImage == "" {
		u.BackgroundImage = defaultConf.BackGroundImage
	}
	return nil
}

// Video 视频实体类
// 普通索引 AuthorId
type Video struct {
	gorm.Model

	Title    string
	AuthorId uint `gorm:"index:idx_author_id"`
	PlayUrl  string
	CoverUrl string
}

// BeforeCreate Hook函数，设置视频默认信息
func (v *Video) BeforeCreate(*gorm.DB) error {
	if v.Title == "" {
		v.Title = defaultConf.VideoTitle
	}
	if v.CoverUrl == "" {
		v.CoverUrl = defaultConf.VideoCover
	}
	return nil
}

//// AuthorVideo 作者-视频关系表
//// 唯一索引 AuthorId -> VideoId
//// 普通索引 VideoId
//type AuthorVideo struct {
//	gorm.Model
//
//	AuthorId int64 `gorm:"uniqueIndex:idx_author_video"`
//	VideoId  int64 `gorm:"uniqueIndex:idx_author_video;index:idx_video"`
//}

// Favorite 点赞表
// 唯一索引 UserId -> VideoId
type Favorite struct {
	gorm.Model

	UserId  int64 `gorm:"uniqueIndex:idx_user_video"`
	VideoId int64 `gorm:"uniqueIndex:idx_user_video"`
}

// Comment 评论表关系表
// 唯一索引 VideoId -> ID(commentID)
type Comment struct {
	VideoId   int64 `gorm:"uniqueIndex:idx_video_comment"`
	ID        uint  `gorm:"uniqueIndex:idx_video_comment;primaryKey"`
	UserId    int64
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// Follow 关注关系表
// 唯一索引 FollowerId -> FollowId
// 普通索引 FollowId
type Follow struct {
	gorm.Model

	FollowerId int64 `gorm:"uniqueIndex:idx_follower_follow"`
	FollowId   int64 `gorm:"uniqueIndex:idx_follower_follow;index:idx_follow"`
}

//// Friend 好友关系表
//// UserId -> FriendId
//type Friend struct {
//	gorm.Model
//	UserId   int64 `gorm:"uniqueIndex:idx_user_friend"`
//	FriendId int64 `gorm:"uniqueIndex:idx_user_friend"`
//}

// Message 用户消息表
// FromUserId -> ToUserId
type Message struct {
	gorm.Model

	FromUserId int64 `gorm:"uniqueIndex:idx_from_to"`
	ToUserId   int64 `gorm:"uniqueIndex:idx_from_to"`
	Content    string
}
