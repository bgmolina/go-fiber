package main

import (
	"github.com/bgmolina/go-fiber/config"
	"github.com/bgmolina/go-fiber/utils"
	"github.com/gofiber/fiber/v2"
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
	logger := utils.LoggerFn("main")
	// load environment variables
	env := config.EnvFn()
	// fiber app instance
	app := fiber.New()
	PORT := env.API_PORT

	//load database connection
	config.DBConnection()

	app.Get("/", func(req *fiber.Ctx) error {
		return req.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Hello 🌎",
		})
	})

	// user routes group
	userGroup := app.Group("/user")
	userGroup.Get("", handlerUser)
	userGroup.Post("", handlerCreateUser)

	logger.Info("Server running at: http://127.0.0.1:" + PORT)
	app.Listen(":" + PORT)
}
