package routing

import (
	"pushpost/internal/services/message_service/transport"
	"pushpost/pkg/middleware"
)

func SetupRoutes(handler transport.MessageHandler) {
	jwtSecret := "bullsonparade"
	messageHandlers := handler.App().Group("/message", middleware.AuthJWTMiddleware(jwtSecret))

	// GET
	messageHandlers.Get("/getByUuid", handler.GetMessagesByUserUUID)

	// POST
	messageHandlers.Post("/create", handler.CreateMessage)

}
