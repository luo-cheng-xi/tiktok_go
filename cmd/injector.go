package cmd

import (
	"tiktok/api"
	"tiktok/internal/middleware"
)

type Injector struct {
	UserController       *api.UserController
	VideoController      *api.VideoController
	LoginCheckMiddleware *middleware.LoginCheckMiddleware
}

func InitInjector(uc *api.UserController, vc *api.VideoController, lcm *middleware.LoginCheckMiddleware) *Injector {
	return &Injector{
		UserController:       uc,
		VideoController:      vc,
		LoginCheckMiddleware: lcm,
	}
}
