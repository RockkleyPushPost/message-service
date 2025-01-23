package transport

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"pushpost/internal/services/message_service/domain"
	"pushpost/internal/services/message_service/domain/dto"
	"pushpost/internal/services/message_service/domain/usecase"
	"pushpost/internal/services/message_service/entity"
)

type MessagesHandler struct {
	useCase domain.MessageUseCase
}

func NewMessagesHandler(useCase usecase.MessageUseCase) *MessagesHandler {
	return &MessagesHandler{useCase: &useCase}
}

func (h *MessagesHandler) CreateMessage(c *fiber.Ctx) error {
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

func (h *MessagesHandler) GetMessagesByUserUUID(c *fiber.Ctx) error {
	var userUUID uuid.UUID

	if err := c.BodyParser(&userUUID); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	//userUUID := c.Context().Value("userUUID")
	//log.Printf("Authorized user UUID: %s", userUUID)
	messages, err := h.useCase.GetMessagesByUserUUID(userUUID)

	if err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(messages)
}
