package services

import (
	"github.com/gofiber/fiber/v2"
)

func GetVersion() fiber.Map {
	return fiber.Map{
		"version": "1.0.0",
	}
}
