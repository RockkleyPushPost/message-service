package usecase

import (
	"github.com/google/uuid"
	"pushpost/internal/services/message_service/domain"
	"pushpost/internal/services/message_service/domain/dto"
	"pushpost/internal/services/message_service/entity"
	"pushpost/internal/services/message_service/storage"
)

// implementation check
var _ domain.MessageUseCase = &MessageUseCase{}

type MessageUseCase struct {
	MessageRepo storage.MessageRepository
}

func (uc *MessageUseCase) CreateMessage(dto *dto.CreateMessageDTO) (err error) {
	message := entity.NewMessage(*dto)

	return uc.MessageRepo.CreateMessage(message)
}

func (uc *MessageUseCase) GetMessagesByUserUUID(uuid uuid.UUID) (messages []entity.Message, err error) {

	return uc.MessageRepo.GetMessagesByUserUUID(uuid)
}
