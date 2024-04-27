package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"

	"github.com/bgmolina/go-fiber/docs"
)

func Routes(app *fiber.App) {
	//enable cors
	app.Use(cors.New())
	//load swagger
	app.Static("/docs", "./docs")
	docs.SwaggerInfo.Title = "Fiber API"
	docs.SwaggerInfo.Description = "API Documentation for Fiber API"

	app.Get("/docs/*", swagger.New(swagger.Config{
		DeepLinking: true,
		URL:         "http://localhost:3000/docs/swagger.yaml",
	}))
	//load all routes
	router := app.Group("/api")
	VersionRouter(router.Group("/version"))
	UserRouter(router.Group("/users"))
}
