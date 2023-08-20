package app

import (
	"log"

	"github.com/biskitsx/Grocery/database"
	"github.com/biskitsx/Grocery/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func SetupAndRunApp() error {
	// dotenv
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	app := fiber.New()
	if err := database.StartMySQL(); err != nil {
		return err
	}

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path} ${latency}\n",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello")
	})
	// routes
	router.SetUpRoutes(app)
	app.Listen(":3001")
	return nil
}
