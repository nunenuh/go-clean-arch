package infrastructure

import (
	"fmt"
	"go-post-clean-arch/internal/auth"
	"go-post-clean-arch/internal/misc"
	"go-post-clean-arch/internal/post"
	"go-post-clean-arch/internal/user"

	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func Run() {
	mariadb, err := ConnectToMariaDB()
	if err != nil {
		log.Fatalf("Database connection failed: %s", err)
	}

	app := fiber.New(fiber.Config{
		AppName:      "Go Post Clean Arch",
		ServerHeader: "Fiber",
	})

	app.Use(cors.New())
	app.Use(compress.New())
	app.Use(etag.New())
	app.Use(favicon.New())
	app.Use(limiter.New(limiter.Config{
		Max: 100,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(&fiber.Map{
				"status":  "fail",
				"message": "Too many requests",
			})
		},
	}))

	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(requestid.New())

	// create repositories
	postRepository := post.NewPostRepository(mariadb)
	userRepository := user.NewUserRepository(mariadb)

	// create all services
	postService := post.NewPostService(postRepository)
	userService := user.NewUserService(userRepository)

	//prepare our endpoints for the API
	misc.NewMiscHandler(app.Group("/api/v1"))
	auth.NewAuthHandler(app.Group("/api/v1/auth"))
	post.NewPostHandler(app.Group("/api/v1/posts"), postService)
	user.NewUserHandler(app.Group("/api/v1/users"), userService)

	app.All("*", func(c *fiber.Ctx) error {
		errorMessage := fmt.Sprintf("Route %s does not exist in this API!", c.OriginalURL())
		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"status":  "fail",
			"message": errorMessage,
		})
	})

	log.Fatal(app.Listen(":8080"))
}
