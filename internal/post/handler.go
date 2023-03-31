package post

import (
	"context"
	"go-post-clean-arch/internal/auth"

	"github.com/gofiber/fiber/v2"
)

type PostHandler struct {
	postService PostService
}

func NewPostHandler(postRoute fiber.Router, ps PostService) {
	handler := &PostHandler{
		postService: ps,
	}

	postRoute.Use(auth.JWTMiddleware(), auth.GetDataFromJWT)

	postRoute.Get("", handler.getPosts)
	postRoute.Post("", handler.createPost)

	postRoute.Get("/:postID", handler.getPost)
	postRoute.Put("/:postID", handler.checkIfPostExistsMiddleware, handler.updatePost)
	postRoute.Delete("/:postID", handler.checkIfPostExistsMiddleware, handler.deletePost)

}

func (h *PostHandler) getPosts(c *fiber.Ctx) error {
	customContex, cancel := context.WithCancel(context.Background())
	defer cancel()

	posts, err := h.postService.FetchPosts(customContex)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   posts,
	})
}

func (h *PostHandler) getPost(c *fiber.Ctx) error {
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	targetedPostID, err := c.ParamsInt("postID")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Please specify a valid post ID",
		})
	}

	post, err := h.postService.FetchPost(customContext, targetedPostID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   post,
	})
}

func (h *PostHandler) createPost(c *fiber.Ctx) error {
	post := &Post{}
	currentUserID := c.Locals("userID").(int)

	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := c.BodyParser(post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	err := h.postService.BuildPost(customContext, post, currentUserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status":  "success",
		"message": "Post created successfully",
	})
}

func (h *PostHandler) updatePost(c *fiber.Ctx) error {
	post := &Post{}
	currentUserID := c.Locals("currentUser").(int)
	targetedPostID := c.Locals("postID").(int)

	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := c.BodyParser(post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	err := h.postService.ModifyPost(customContext, targetedPostID, post, currentUserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Post updated successfully",
	})

}

func (h *PostHandler) deletePost(c *fiber.Ctx) error {
	targetedPostID := c.Locals("postID").(int)

	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := h.postService.DestroyPost(customContext, targetedPostID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
