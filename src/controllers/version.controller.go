package controllers

import (
	"github.com/bgmolina/go-fiber/src/services"
	"github.com/gofiber/fiber/v2"
)

func GetVersion(req *fiber.Ctx) error {
	response := services.GetVersion()
	return req.Status(fiber.StatusOK).JSON(response)
}
