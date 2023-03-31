package post

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

func (h *PostHandler) checkIfPostExistsMiddleware(c *fiber.Ctx) error {
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	targetedPostID, err := c.ParamsInt("postID")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	searchedPost, err := h.postService.FetchPost(customContext, targetedPostID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	if searchedPost == nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "There is no post with this ID!",
		})
	}

	c.Locals("postID", targetedPostID)
	return c.Next()

}
