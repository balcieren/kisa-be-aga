package v1

import (
	"backend/app/url-api/v1/handler"
	"backend/app/url-api/v1/service"

	"go.uber.org/fx"
)

var Module = fx.Provide(
	service.New,
	handler.New,
	NewRouter,
)
