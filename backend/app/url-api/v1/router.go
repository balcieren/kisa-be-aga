package v1

import (
	"backend/app/url-api/v1/handler"

	"github.com/gofiber/fiber/v2"
)

type Router struct {
	url     fiber.Router
	handler *handler.Handler
}

func NewRouter(app *fiber.App, h *handler.Handler) *Router {
	return &Router{
		url:     app.Group("/v1/urls"),
		handler: h,
	}
}

func (r *Router) Setup() {
	r.url.Get("/:short_id", r.handler.RedirectToURL)
	r.url.Post("/", r.handler.ShortenURL)
}
