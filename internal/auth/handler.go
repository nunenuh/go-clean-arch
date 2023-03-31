package auth

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type AuthHandler struct{}

func NewAuthHandler(authRoute fiber.Router) {
	handler := &AuthHandler{}

	authRoute.Post("/login", handler.signInUser)
	authRoute.Post("/logout", handler.signOutUser)
	authRoute.Get("private", JWTMiddleware(), handler.privateRoute)

}

func (h *AuthHandler) signInUser(c *fiber.Ctx) error {
	type loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	type jwtClaims struct {
		UserID string `json:"user_id"`
		User   string `json:"user"`
		jwt.StandardClaims
	}

	request := &loginRequest{}
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	if request.Username != os.Getenv("API_USERNAME") || request.Password != os.Getenv("API_PASSWORD") {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Invalid username or password",
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwtClaims{
		os.Getenv("API_USERID"),
		os.Getenv("API_USERNAME"),
		jwt.StandardClaims{
			Audience:  "go-post-clean-arch-users",
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "go-post-clean-arch",
		},
	})

	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    signedToken,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 24),
		Secure:   false,
		HTTPOnly: true,
	})

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"token":  signedToken,
	})
}

func (h *AuthHandler) signOutUser(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    "loggedOut",
		Path:     "/",
		Expires:  time.Now().Add(-time.Hour * 24),
		Secure:   false,
		HTTPOnly: true,
	})

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Logged out succesfully",
	})
}

func (h *AuthHandler) privateRoute(c *fiber.Ctx) error {
	type jwtResponse struct {
		UserID interface{} `json:"user_id"`
		User   interface{} `json:"user"`
		Iss    interface{} `json:"iss"`
		Aud    interface{} `json:"aud"`
		Exp    interface{} `json:"exp"`
	}

	jwtData := c.Locals("user").(*jwt.Token)
	claims := jwtData.Claims.(jwt.MapClaims)

	jwtResp := &jwtResponse{
		UserID: claims["uid"],
		User:   claims["user"],
		Iss:    claims["iss"],
		Aud:    claims["aud"],
		Exp:    claims["exp"],
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Welcome to the private route!",
		"jwtData": jwtResp,
	})
}
