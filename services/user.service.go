package services

import (
	"github.com/bgmolina/go-fiber/structs"
	"github.com/bgmolina/go-fiber/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var logger = utils.GetLogger("service")

func GetUser() structs.User {
	user := structs.User{
		Name:     "Moni",
		LastName: "Argento",
		Age:      35,
	}
	return user
}

func CreateUser(req *fiber.Ctx) structs.User {
	user := new(structs.User)
	err := req.BodyParser(user)
	if err != nil {
		logger.Error("Error parsing body")
		//TODO: pasar error por middleware [!]
	}
	user.Id = uuid.NewString() // generate a new UUID
	return *user
}
