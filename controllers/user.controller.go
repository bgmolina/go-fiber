package controllers

import (
	"github.com/bgmolina/go-fiber/services"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(req *fiber.Ctx) error {
	response := services.GetUsers()
	return req.Status(fiber.StatusOK).JSON(response)
}

func GetUserById(req *fiber.Ctx) error {
	id := req.Params("id")
	response := services.GetUserById(id)
	return req.Status(fiber.StatusOK).JSON(response)
}

func CreateUser(req *fiber.Ctx) error {
	response := services.CreateUser(req)
	return req.Status(fiber.StatusCreated).JSON(response)
}

func DeleteUserById(req *fiber.Ctx) error {
	id := req.Params("id")
	response := services.DeleteUserById(id)
	return req.Status(fiber.StatusNoContent).JSON(response)
}
