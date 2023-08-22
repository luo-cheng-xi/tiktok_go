package model

type UserVO struct {
	ID              uint64 `json:"id"`
	Name            string `json:"name"`
	FollowCount     uint64 `json:"follow_count"`
	FollowerCount   uint64 `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	TotalFavorited  uint64 `json:"total_favorited"`
	WorkCount       uint64 `json:"work_count"`
	FavoriteCount   uint64 `json:"favorite_count"`
}

type VideoVO struct {
	ID            uint64 `json:"id"`
	Author        UserVO `json:"author"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount uint64 `json:"favorite_count"`
	CommentCount  uint64 `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
	Title         string `json:"title"`
}

type MessageVO struct {
	ID         uint64 `json:"id"`
	ToUserId   uint64 `json:"to_user_id"`
	FromUserId uint64 `json:"from_user_id"`
	Content    string `json:"content"`
	CreateTime string `json:"create_time"`
}
