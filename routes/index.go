package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
)

func Routes(app *fiber.App) {
	//enable cors
	app.Use(cors.New())
	//load swagger
	app.Static("/docs", "./docs", fiber.Static{
		CacheDuration: 0,
	})

	app.Get("/docs/*", swagger.New(swagger.Config{
		DeepLinking: true,
		URL:         "http://localhost:3000/docs/swagger.yaml",
		Title:       "Documentation",
	}))

	//load all routes
	router := app.Group("/api")
	VersionRouter(router.Group("/version"))
	UserRouter(router.Group("/user"))
}
