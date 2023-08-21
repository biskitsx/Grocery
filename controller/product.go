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
	AddOrRemoveToCart(c *fiber.Ctx) error
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

func (controller *productController) AddOrRemoveToCart(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return fiber.NewError(400, "product not found")
	}

	// Find the product
	product := model.Product{}
	if err := database.Db.Where("id = ?", id).First(&product).Error; err != nil {
		return fiber.NewError(400, "product not found")
	}

	// Find the user
	user := model.User{}
	userId := c.Locals("userId") // Assuming userId is of type uint
	if err := database.Db.Where("id = ?", userId).Preload("Cart").First(&user).Error; err != nil {
		return fiber.NewError(400, "user not found")
	}

	// Check if the product is already in the cart
	productFoundInCart := false
	for _, cartProduct := range user.Cart {
		if cartProduct == product {
			productFoundInCart = true
			break
		}
	}

	if productFoundInCart {
		database.Db.Preload("Cart").Model(&user).Association("Cart").Delete(&product)
	} else {
		database.Db.Model(&user).Association("Cart").Append(&product)
	}
	return c.JSON(user)
}
