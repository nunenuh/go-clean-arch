package user

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService UserService
}

func NewUserHandler(userRoute fiber.Router, us UserService) {
	handler := &UserHandler{
		userService: us,
	}

	userRoute.Get("", handler.getUsers)
	userRoute.Post("", handler.createUser)

	userRoute.Get("/:userID", handler.getUser)
	userRoute.Put("/:userID", handler.checkIfUserExistsMiddleware, handler.updateUser)
	userRoute.Delete("/:userID", handler.checkIfUserExistsMiddleware, handler.deleteUser)

}

func (h *UserHandler) getUsers(c *fiber.Ctx) error {
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	users, err := h.userService.GetUsers(customContext)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   users,
	})
}

func (h *UserHandler) getUser(c *fiber.Ctx) error {
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	targetedUserID, err := c.ParamsInt("userID")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Please specify a valid user ID",
		})
	}

	user, err := h.userService.GetUser(customContext, targetedUserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   user,
	})
}

func (h *UserHandler) createUser(c *fiber.Ctx) error {
	user := &User{}

	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	err := h.userService.CreateUser(customContext, user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status": "success",
		"data":   user,
	})
}

func (h *UserHandler) updateUser(c *fiber.Ctx) error {
	user := &User{}
	targetedUserID := c.Locals("userID").(int)

	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	err := h.userService.UpdateUser(customContext, targetedUserID, user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   user,
	})
}

func (h *UserHandler) deleteUser(c *fiber.Ctx) error {
	targetedUserID := c.Locals("userID").(int)

	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := h.userService.DeleteUser(customContext, targetedUserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
