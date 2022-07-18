package handler

import (
	"backend/internal/url-api/v1/service"
	"backend/pkg/config"
)

type Handler struct {
	service *service.Service
	config  *config.Config
}

func New(s *service.Service, c *config.Config) *Handler {
	return &Handler{
		service: s,
		config:  c,
	}
}
