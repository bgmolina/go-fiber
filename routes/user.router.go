package routes

import (
	"github.com/bgmolina/go-fiber/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(router fiber.Router) {
	router.Get("", controllers.GetUsers)
	router.Get(":id", controllers.GetUserById)
	router.Post("", controllers.CreateUser)
	router.Delete(":id", controllers.DeleteUserById)
}
