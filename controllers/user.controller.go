package controllers

import (
	"github.com/bgmolina/go-fiber/services"
	"github.com/gofiber/fiber/v2"
)

type IUserController interface {
	GetUser(req *fiber.Ctx) error
}

type UserController struct{}

func GetUser(req *fiber.Ctx) error {
	response := services.GetUser()
	return req.Status(fiber.StatusOK).JSON(response)
}

func CreateUser(req *fiber.Ctx) error {
	response := services.CreateUser(req)
	return req.Status(fiber.StatusCreated).JSON(response)
}
