package model

type UserVO struct {
	ID              uint   `json:"id"`
	Name            string `json:"name"`
	FollowCount     int64  `json:"follow_count"`
	FollowerCount   int64  `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	TotalFavorited  int64  `json:"total_favorited"`
	WorkCount       int64  `json:"work_count"`
	FavoriteCount   int64  `json:"favorite_count"`
}

// ParseUserVO 待补全
func ParseUserVO(user User) UserVO {
	return UserVO{
		ID:              user.ID,
		Name:            user.Username,
		Avatar:          user.Avatar,
		BackgroundImage: user.BackgroundImage,
		Signature:       user.Signature,
	}
}

type VideoVO struct {
	ID            uint   `json:"id"`
	Author        UserVO `json:"author"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
	Title         string `json:"title"`
}

// ParseVideoVO 待补全
func ParseVideoVO(video Video) VideoVO {
	return VideoVO{
		ID: video.ID,
		//Author:
		PlayUrl:  video.PlayUrl,
		CoverUrl: video.CoverUrl,
		//FavoriteCount:
		//CommentCount
		//IsFavorite:
		Title: video.Title,
	}
}

type MessageVO struct {
	ID         uint   `json:"id"`
	ToUserId   int64  `json:"to_user_id"`
	FromUserId int64  `json:"from_user_id"`
	Content    string `json:"content"`
	CreateTime string `json:"create_time"`
}
