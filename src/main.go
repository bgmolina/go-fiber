package main

import (
	"github.com/bgmolina/go-fiber/src/config"
	"github.com/bgmolina/go-fiber/src/routes"
	"github.com/bgmolina/go-fiber/src/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	logger := utils.GetLogger("main")
	env := config.EnvFn()
	app := fiber.New()
	PORT := env.API_PORT

	//load routes
	routes.Routes(app)

	logger.Info("Server running at: http://127.0.0.1:" + PORT)
	app.Listen(":" + PORT)
}
