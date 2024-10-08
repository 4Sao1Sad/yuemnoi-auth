package notification

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sds-2/config"
)

type NofiticationHandler struct {
	cfg *config.Config
}

func NewNofiticationHandler() *NofiticationHandler {
	return &NofiticationHandler{}
}

func (h *NofiticationHandler) PushNotification(c *fiber.Ctx) error {
	return c.SendStatus(200)
}
