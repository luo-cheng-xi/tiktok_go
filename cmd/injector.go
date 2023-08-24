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
	LoginCheckMiddleware *middleware.LoginCheckMiddleware
}

func InitInjector(
	uc *api.UserController,
	vc *api.VideoController,
	rc *api.RelationController,
	ic *api.FavoriteController,
	lcm *middleware.LoginCheckMiddleware) *Injector {
	return &Injector{
		UserController:       uc,
		VideoController:      vc,
		relationController:   rc,
		favoriteController:   ic,
		LoginCheckMiddleware: lcm,
	}
}
