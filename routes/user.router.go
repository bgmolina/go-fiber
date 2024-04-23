package routes

import (
	"github.com/bgmolina/go-fiber/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(router fiber.Router) {
	router.Get("/", controllers.GetUser)
	router.Post("/", controllers.CreateUser)
}
