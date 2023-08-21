package middleware

import (
	"github.com/biskitsx/Grocery/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func AuthVerify(c *fiber.Ctx) error {
	token := c.Cookies("access_token")
	if token == "" {
		return fiber.NewError(405, "unauthorize")
	}
	validatedToken, err := service.NewJwtService().ValidateToken(token)
	if err != nil {
		return fiber.NewError(fiber.StatusNonAuthoritativeInformation, "user not authenticated")
	}
	payload, _ := validatedToken.Claims.(jwt.MapClaims)
	id, ok := payload["id"].(float64)
	if !ok {
		return fiber.NewError(fiber.StatusConflict, "int error")
	}
	c.Locals("userId", id)
	return c.Next()
}
