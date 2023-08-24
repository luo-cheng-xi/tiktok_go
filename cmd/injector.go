package cmd

import (
	"tiktok/api"
	"tiktok/internal/middleware"
)

type Injector struct {
	UserController       *api.UserController
	VideoController      *api.VideoController
	relationController   *api.RelationController
	favoriteController   *api.FavoriteController
	commentController    *api.CommentController
	LoginCheckMiddleware *middleware.LoginCheckMiddleware
}

func InitInjector(
	uc *api.UserController,
	vc *api.VideoController,
	rc *api.RelationController,
	cc *api.CommentController,
	ic *api.FavoriteController,
	lcm *middleware.LoginCheckMiddleware) *Injector {
	return &Injector{
		UserController:       uc,
		VideoController:      vc,
		relationController:   rc,
		commentController:    cc,
		favoriteController:   ic,
		LoginCheckMiddleware: lcm,
	}
}
