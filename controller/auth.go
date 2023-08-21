package controller

import (
	"github.com/biskitsx/Grocery/database"
	"github.com/biskitsx/Grocery/model"
	"github.com/biskitsx/Grocery/service"
	"github.com/gofiber/fiber/v2"
)

type AuthController interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}

type authController struct {
	jwtService service.JwtService
}

func NewAuthController() AuthController {
	return &authController{
		jwtService: service.NewJwtService(),
	}
}

type RegisterDTO struct {
	Username string
	Password string
}

type LoginDTO struct {
	Username string
	Password string
}

func (controller *authController) Register(c *fiber.Ctx) error {
	dto := &RegisterDTO{}
	if err := c.BodyParser(dto); err != nil {
		return fiber.NewError(400, "All field is required")
	}
	user := &model.User{}
	user.Username = dto.Username
	user.Password = dto.Password
	database.Db.Create(user)
	return c.JSON(user)
}

func (controller *authController) Login(c *fiber.Ctx) error {
	dto := &LoginDTO{}
	if err := c.BodyParser(dto); err != nil {
		return fiber.NewError(400, "All field is required")
	}
	user := &model.User{}
	database.Db.Where("Username = ?", dto.Username).Preload("Cart").First(user)
	if user.Username == "" {
		return fiber.NewError(405, "Username not found")
	}

	if user.Password != dto.Password {
		return fiber.NewError(405, "Wrong Password")
	}
	token := controller.jwtService.GenerateToken(user.ID)
	cookieToken := &fiber.Cookie{
		Name:     "access_token",
		Value:    token,
		Secure:   false,
		HTTPOnly: true,
	}

	c.Cookie(cookieToken)
	return c.JSON(user)
}
