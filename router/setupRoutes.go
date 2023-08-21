package router

import (
	"github.com/biskitsx/Grocery/controller"
	"github.com/biskitsx/Grocery/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	productRoutes(app)
	categoryRoutes(app)
	authRoutes(app)
	userRoutes(app)
}

func productRoutes(app *fiber.App) {
	product := controller.NewProductController()
	app.Post("/api/product", product.CreateProduct)
	app.Get("/api/product", product.GetProduct)
	app.Delete("/api/product/:id", product.GetProduct)
	app.Put("/api/product/add-remove/:id", middleware.AuthVerify, product.AddOrRemoveToCart)
}

func categoryRoutes(app *fiber.App) {
	category := controller.NewCategoryController()
	app.Post("/api/category", category.CreateCategory)
	app.Get("/api/category", middleware.AuthVerify, category.GetCategory)
	app.Get("/api/category/:id", category.GetCategoryById)
}

func authRoutes(app *fiber.App) {
	auth := controller.NewAuthController()
	app.Post("/api/auth/login", auth.Login)
	app.Post("/api/auth/register", auth.Register)
}

func userRoutes(app *fiber.App) {
	user := controller.NewUserController()
	app.Get("/api/user", user.GetUser)
	app.Get("/api/user/cart", middleware.AuthVerify, user.GetUserCart)
}
