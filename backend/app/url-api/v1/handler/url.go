package handler

import (
	"backend/pkg/ent"
	"backend/pkg/proto"
	"context"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) RedirectToURL(c *fiber.Ctx) error {
	sid := c.Params("short_id")

	u, err := h.service.Url.Get(context.Background(), &proto.GetUrlRequest{
		ShortId: sid,
	})
	if ent.IsNotFound(err) {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.Redirect(u.GetUrl())
}

func (h *Handler) ShortenURL(c *fiber.Ctx) error {
	body := new(proto.CreateUrlRequest)
	if err := c.BodyParser(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	url, err := h.service.Url.Create(context.Background(), &proto.CreateUrlRequest{
		Url: body.Url,
	})
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(url)
}
