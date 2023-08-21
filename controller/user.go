package controller

import (
	"github.com/biskitsx/Grocery/database"
	"github.com/biskitsx/Grocery/model"
	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	GetUser(c *fiber.Ctx) error
	GetUserCart(c *fiber.Ctx) error
}
type userController struct {
}

func NewUserController() UserController {
	return &userController{}
}

func (controller *userController) GetUser(c *fiber.Ctx) error {
	user := new([]model.User)
	database.Db.Preload("Cart").Find(user)
	return c.JSON(user)
}

func (controller *userController) GetUserCart(c *fiber.Ctx) error {
	id := c.Locals("userId")
	user := new(model.User)
	database.Db.Preload("Cart").Where("id = ?", id).First(user)
	cart := user.Cart
	return c.JSON(cart)
}
