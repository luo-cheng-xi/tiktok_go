//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/google/wire"
	"tiktok/api"
	"tiktok/internal/conf"
	"tiktok/internal/data"
	"tiktok/internal/middleware"
	"tiktok/internal/service"
	"tiktok/pkg/logging"
	"tiktok/pkg/util"
)

func BuildInjector() (*Injector, error) {
	wire.Build(
		InitInjector,
		api.NewUserController,
		api.NewVideoController,
		middleware.NewLoginCheck,
		service.NewUserService,

		data.ProviderSet,
		conf.ProviderSet,
		util.ProviderSet,

		logging.NewLogger,
	)
	return &Injector{}, nil
}
