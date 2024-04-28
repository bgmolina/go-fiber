package services

import (
	"github.com/bgmolina/go-fiber/src/config"
	"github.com/gofiber/fiber/v2"
)

func GetVersion() fiber.Map {
	env := config.EnvFn()
	return fiber.Map{
		"version": env.API_VERSION,
	}
}
