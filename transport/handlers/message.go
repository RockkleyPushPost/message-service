package transport

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"pushpost/internal/services/message_service/domain"
	"pushpost/internal/services/message_service/domain/dto"
	"pushpost/internal/services/message_service/entity"
)

type MessageHandler struct {
	useCase domain.MessageUseCase
	app     *fiber.App
}

func NewMessagesHandler(useCase domain.MessageUseCase, app *fiber.App) *MessageHandler {
	return &MessageHandler{useCase: useCase, app: app}
}

func (h *MessageHandler) CreateMessage(c *fiber.Ctx) error {
	var body entity.Message

	if err := c.BodyParser(&body); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	params := dto.CreateMessageDTO{
		SenderUUID:   body.SenderUUID,
		ReceiverUUID: body.ReceiverUUID,
		Content:      body.Content,
	}

	if err := params.Validate(); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.useCase.CreateMessage(&params)

	if err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Message created successfully"})
}

func (h *MessageHandler) GetMessagesByUserUUID(c *fiber.Ctx) error {
	var userUUID uuid.UUID

	if err := c.BodyParser(&userUUID); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	messages, err := h.useCase.GetMessagesByUserUUID(userUUID)

	if err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(messages)
}

func (h *MessageHandler) App() *fiber.App {
	return h.app
}
