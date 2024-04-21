package main

import (
	"github.com/bgmolina/go-fiber/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/google/uuid"
)

type User struct {
	Id       string
	Name     string
	LastName string
	Age      int
}

func handlerUser(req *fiber.Ctx) error {
	user := User{
		Name:     "Moni",
		LastName: "Argento",
		Age:      35,
	}
	return req.Status(fiber.StatusOK).JSON(user)
}

func handlerCreateUser(req *fiber.Ctx) error {
	user := new(User)
	// if body parsing fails, return an error
	if err := req.BodyParser(user); err != nil {
		return err
	}
	user.Id = uuid.NewString() // generate a new UUID
	return req.Status(fiber.StatusCreated).JSON(user)
}

func main() {
	// fiber app instance
	app := fiber.New()
	PORT := config.Config("PORT")

	//middleware
	// log all requests
	app.Use(logger.New())

	app.Get("/", func(req *fiber.Ctx) error {
		return req.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Hello 🌎",
		})
	})

	// user routes group
	userGroup := app.Group("/user")
	userGroup.Get("", handlerUser)
	userGroup.Post("", handlerCreateUser)

	app.Listen(":" + PORT)
}
