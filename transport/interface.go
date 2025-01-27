package transport

import "github.com/gofiber/fiber/v2"

type MessageHandler interface {
	CreateMessage(c *fiber.Ctx) error
	GetMessagesByUserUUID(c *fiber.Ctx) error
	App() *fiber.App
}
