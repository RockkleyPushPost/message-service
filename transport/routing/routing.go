package routing

import (
	"github.com/gofiber/fiber/v2"
	"pushpost/internal/di"
	"pushpost/pkg/middleware"
)

func SetupRoutes(app *fiber.App, container di.Container) {
	jwtSecret := "shenanigans"
	messageHandlers := app.Group("/message", middleware.AuthJWTMiddleware(jwtSecret))

	// GET
	messageHandlers.Get("/getByUuid", container.MessageHandler.GetMessagesByUserUUID)

	// POST
	messageHandlers.Post("/create", container.MessageHandler.CreateMessage)

}
