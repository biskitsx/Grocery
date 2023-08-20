package router

import (
	"github.com/biskitsx/Grocery/controller"
	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	productRoutes(app)
	categoryRoutes(app)
	authRoutes(app)
}

func productRoutes(app *fiber.App) {
	product := controller.NewProductController()
	app.Post("/api/product", product.CreateProduct)
	app.Get("/api/product", product.GetProduct)
	app.Delete("/api/product/id", product.GetProduct)
}

func categoryRoutes(app *fiber.App) {
	category := controller.NewCategoryController()
	app.Post("/api/category", category.CreateCategory)
	app.Get("/api/category", category.GetCategory)
	app.Get("/api/category/:id", category.GetCategoryById)
}

func authRoutes(app *fiber.App) {
	auth := controller.NewAuthController()
	app.Post("/api/auth/login", auth.Login)
	app.Post("/api/auth/register", auth.Register)
}
