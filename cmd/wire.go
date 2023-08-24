//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/google/wire"
	"tiktok/api"
	"tiktok/internal/conf"
	"tiktok/internal/data"
	"tiktok/internal/manager"
	"tiktok/internal/middleware"
	"tiktok/internal/service"
	"tiktok/pkg/logging"
	"tiktok/pkg/util"
)

func BuildInjector() (*Injector, error) {
	wire.Build(
		InitInjector,
		middleware.NewLoginCheck,

		api.ProviderSet,

		service.ProviderSet,

		manager.ProviderSet,

		data.ProviderSet,

		conf.ProviderSet,

		util.ProviderSet,

		logging.NewLogger,
	)
	return &Injector{}, nil
}
