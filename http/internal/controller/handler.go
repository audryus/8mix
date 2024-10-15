package controller

import (
	"github.com/audryus/8mix/http/internal/usecase"
	"github.com/audryus/8mix/http/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	uc     *usecase.PlaylistUC
	logger *logger.Log
}

func New(app *fiber.App, uc *usecase.PlaylistUC, logger *logger.Log) {
	r := &Controller{
		uc:     uc,
		logger: logger,
	}

	app.Get("/", r.Index)
}

func (h *Controller) Index(c *fiber.Ctx) error {
	h.logger.Info("index")
	return c.Render("index", fiber.Map{})
}
