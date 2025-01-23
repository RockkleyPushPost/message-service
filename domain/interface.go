package domain

import (
	"github.com/google/uuid"
	"pushpost/internal/services/message_service/domain/dto"
	"pushpost/internal/services/message_service/entity"
)

type MessageUseCase interface {
	CreateMessage(dto *dto.CreateMessageDTO) (err error)
	GetMessagesByUserUUID(uuid uuid.UUID) (messages []entity.Message, err error)
}
