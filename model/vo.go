package model

type UserVO struct {
	Id              int64
	Name            string
	FollowCount     int64
	FollowerCount   int64
	IsFollow        bool
	Avatar          string
	BackgroundImage string
	Signature       string
	TotalFavorited  int64
	WorkCount       int64
	FavoriteCount   int64
}
