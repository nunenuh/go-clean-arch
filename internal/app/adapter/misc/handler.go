package misc

import "github.com/gofiber/fiber/v2"

type MiscHandler struct{}

func NewMiscHandler(miscRoute fiber.Router) {
	handler := &MiscHandler{}

	miscRoute.Get("", handler.healtCheck)
}

func (h *MiscHandler) healtCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   "OK",
	})
}
