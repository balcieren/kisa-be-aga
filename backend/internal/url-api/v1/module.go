package v1

import (
	"backend/internal/url-api/v1/handler"
	"backend/internal/url-api/v1/service"

	"go.uber.org/fx"
)

var Module = fx.Provide(
	service.New,
	handler.New,
	NewRouter,
)
