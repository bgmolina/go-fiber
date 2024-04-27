package routes

import (
	"github.com/bgmolina/go-fiber/controllers"
	"github.com/gofiber/fiber/v2"
)

func VersionRouter(router fiber.Router) {
	router.Get("", controllers.GetVersion)
}
