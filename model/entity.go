package model

import (
	"gorm.io/gorm"
	"time"
)

// User 用户实体类
// 唯一联合索引 Username -> Password
type User struct {
	Username        string `gorm:"size:32;uniqueIndex:idx_username_password"`
	Password        string `gorm:"size:32;uniqueIndex:idx_username_password"`
	Avatar          string `gorm:"size:255"`
	BackgroundImage string `gorm:"size:255"`
	Signature       string `gorm:"size:200"`
	gorm.Model
}

// Video 视频实体类
// 普通索引 AuthorId
type Video struct {
	Title    string
	AuthorId string `gorm:"index:idx_author_id"`
	PlayUrl  string
	CoverUrl string
	gorm.Model
}

// AuthorVideo 作者-视频关系表
// 唯一索引 AuthorId -> VideoId
// 普通索引 VideoId
type AuthorVideo struct {
	AuthorId int64 `gorm:"uniqueIndex:idx_author_video"`
	VideoId  int64 `gorm:"uniqueIndex:idx_author_video;index:idx_video"`
	gorm.Model
}

// Favorite 点赞表
// 唯一索引 UserId -> VideoId
type Favorite struct {
	UserId  int64 `gorm:"uniqueIndex:idx_user_video"`
	VideoId int64 `gorm:"uniqueIndex:idx_user_video"`
	gorm.Model
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
	FollowerId int64 `gorm:"uniqueIndex:idx_follower_follow"`
	FollowId   int64 `gorm:"uniqueIndex:idx_follower_follow;index:idx_follow"`
	gorm.Model
}

// Friend 好友关系表
// UserId -> FriendId
type Friend struct {
	UserId   int64 `gorm:"uniqueIndex:idx_user_friend"`
	FriendId int64 `gorm:"uniqueIndex:idx_user_friend"`
	gorm.Model
}

// Message 用户消息表
// FromUserId -> ToUserId
type Message struct {
	FromUserId int64 `gorm:"uniqueIndex:idx_from_to"`
	ToUserId   int64 `gorm:"uniqueIndex:idx_from_to"`
	Content    string
	gorm.Model
}
