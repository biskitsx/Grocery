package controller

import (
	"github.com/biskitsx/Grocery/database"
	"github.com/biskitsx/Grocery/model"
	"github.com/biskitsx/Grocery/service"
	"github.com/gofiber/fiber/v2"
)

type ProductController interface {
	CreateProduct(c *fiber.Ctx) error
	GetProduct(c *fiber.Ctx) error
	DeleteProductById(c *fiber.Ctx) error
}

type productController struct {
	cldService service.CloudinaryService
}

func NewProductController() ProductController {
	return &productController{
		cldService: service.NewCloudinaryService(),
	}
}

type ProductDTO struct {
	Name       string
	Price      uint
	CategoryID uint
}

func (controller *productController) CreateProduct(c *fiber.Ctx) error {
	dto := &ProductDTO{}
	if err := c.BodyParser(dto); err != nil {
		return fiber.NewError(400, err.Error())
	}
	file, _ := c.FormFile("upload")
	imageUrl, _ := controller.cldService.UploadImage(file)
	product := &model.Product{
		Name:       dto.Name,
		CategoryID: dto.CategoryID,
		Price:      dto.Price,
		Picture:    imageUrl,
	}

	database.Db.Create(product)
	return c.JSON(product)
}

func (controller *productController) GetProduct(c *fiber.Ctx) error {
	product := &[]model.Product{}
	database.Db.Find(product)
	return c.JSON(product)
}

func (controller *productController) DeleteProductById(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return fiber.NewError(400, "id not found")
	}
	product := &model.Product{}
	database.Db.Delete(product, id)
	return c.JSON(fiber.Map{
		"msg": "Delete successfully",
	})
}
