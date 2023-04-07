package auth

import (
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt"
)

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "error",
			"message": "Missing or malformed JWT",
		})
	}

	return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
		"status":  "error",
		"message": "Invalid or expired JWT",
	})
}

func JWTMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		ErrorHandler:  jwtError,
		SigningKey:    []byte(os.Getenv("JWT_SECRET")),
		SigningMethod: "HS256",
		TokenLookup:   "cookie:jwt",
	})
}

func GetDataFromJWT(c *fiber.Ctx) error {
	jwtData := c.Locals("user").(*jwt.Token)
	claims := jwtData.Claims.(jwt.MapClaims)
	parsedUserID := claims["uid"].(string)
	userID, err := strconv.Atoi(parsedUserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	c.Locals("currentUser", userID)
	return c.Next()

}
