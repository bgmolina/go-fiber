package routes

import "github.com/gofiber/fiber/v2"

func version(req *fiber.Ctx) error {
	return req.Status(fiber.StatusOK).JSON(fiber.Map{
		"version": "1.0.0",
	})
}

func Routes(app *fiber.App) {
	router := app.Group("/api")
	router.Get("", version)
	UserRouter(router.Group("/user"))
}
