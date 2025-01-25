package config

import (
	"github.com/gofiber/fiber/v2"
	"pushpost/pkg/database"
)

type Config struct {
	Database database.Config
	Fiber    fiber.Config
}
