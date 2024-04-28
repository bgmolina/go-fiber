package routes

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/bgmolina/go-fiber/src/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
)

func Routes(app *fiber.App) {
	env := config.EnvFn()
	pathDocs := func() string {
		var result string
		if env.ENV == "DEV" {
			result = "./src/docs"
		}
		if env.ENV == "PROD" {
			baseDir, _ := os.Getwd()
			result = filepath.Join(baseDir, "docs")
		}
		return result
	}()

	//enable cors
	app.Use(cors.New())

	//load swagger
	app.Static("/docs", pathDocs, fiber.Static{
		CacheDuration: 0,
	})

	app.Get("/docs/*", swagger.New(swagger.Config{
		DeepLinking: true,
		URL:         fmt.Sprintf("%s:%s/docs/swagger.yaml", env.API_HOST, env.API_PORT),
		Title:       "Documentation",
	}))

	//load all routes
	router := app.Group("/api")
	VersionRouter(router.Group("/version"))
	UserRouter(router.Group("/user"))
}
