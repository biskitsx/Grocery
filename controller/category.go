package controller

import (
	"github.com/biskitsx/Grocery/database"
	"github.com/biskitsx/Grocery/model"
	"github.com/gofiber/fiber/v2"
)

type CategoryController interface {
	GetCategory(c *fiber.Ctx) error
	CreateCategory(c *fiber.Ctx) error
	GetCategoryById(c *fiber.Ctx) error
}
type categoryController struct {
}

func NewCategoryController() CategoryController {
	return &categoryController{}
}

type CreateCategoryDTO struct {
	Name string
}

func (controller *categoryController) CreateCategory(c *fiber.Ctx) error {
	dto := &CreateCategoryDTO{}
	if err := c.BodyParser(dto); err != nil {
		return fiber.NewError(404, "error")
	}
	category := &model.Category{}
	category.Name = dto.Name
	database.Db.Create(category)
	return c.JSON(category)
}

func (controller *categoryController) GetCategory(c *fiber.Ctx) error {
	category := &[]model.Category{}
	database.Db.Preload("Products").Find(category)
	return c.JSON(category)
}

func (controller *categoryController) GetCategoryById(c *fiber.Ctx) error {
	id := c.Params("id")
	category := &model.Category{}
	database.Db.Select("name").Where("id = ?", id).First(category)
	return c.JSON(category)
}
