package services

import (
	"github.com/bgmolina/go-fiber/structs"
	"github.com/bgmolina/go-fiber/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var logger = utils.GetLogger("service")

func GetUsers() []structs.User {
	return []structs.User{
		{
			Id:       uuid.NewString(),
			Name:     "Coqui",
			LastName: "Argento",
			Age:      25,
		},
		{
			Id:       uuid.NewString(),
			Name:     "Paola",
			LastName: "Argento",
			Age:      30,
		},
	}
}

func GetUserById(id string) structs.User {
	return structs.User{
		Id:       id,
		Name:     "Moni",
		LastName: "Argento",
		Age:      35,
	}
}

func CreateUser(req *fiber.Ctx) structs.User {
	user := new(structs.User)
	err := req.BodyParser(user)
	if err != nil {
		logger.Error("Error parsing body")
	}
	user.Id = uuid.NewString()
	return *user
}

func DeleteUserById(id string) fiber.Map {
	return fiber.Map{
		"message": "User deleted: " + id,
	}
}
