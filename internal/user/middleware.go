package user

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

func (h *UserHandler) checkIfUserExistsMiddleware(c *fiber.Ctx) error {
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	targetedUserID, err := c.ParamsInt("userID")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	searchedUser, err := h.userService.GetUser(customContext, targetedUserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	if searchedUser == nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "There is no user with this ID!",
		})
	}

	c.Locals("userID", targetedUserID)
	return c.Next()
}
