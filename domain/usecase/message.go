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
	if err = dto.Validate(); err != nil {
		return
	}

	message := entity.Message{
		UUID:         uuid.New(),
		SenderUUID:   dto.SenderUUID,
		ReceiverUUID: dto.ReceiverUUID,
		Content:      dto.Content,
	}
	if err = uc.MessageRepo.CreateMessage(&message); err != nil {

		return err
	}

	return
}

func (uc *MessageUseCase) GetMessagesByUserUUID(uuid uuid.UUID) (messages []entity.Message, err error) {
	messages, err = uc.MessageRepo.GetMessagesByUserUUID(uuid)

	if err != nil {

		return
	}

	return
}
