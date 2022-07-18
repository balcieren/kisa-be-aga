package main

import (
	urlapi "backend/internal/url-api"
	v1 "backend/internal/url-api/v1"
	"backend/pkg/config"
	"backend/pkg/logger"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(config.New),
		fx.Provide(logger.New),
		fx.Provide(urlapi.New),
		fx.Module("v1", v1.Module),
		fx.Invoke(urlapi.Start),
	).Run()
}
