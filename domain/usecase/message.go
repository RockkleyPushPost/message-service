package usecase

import (
	"github.com/google/uuid"
	"pushpost/internal/services/message_service/domain/dto"
	"pushpost/internal/services/message_service/entity"
	"pushpost/internal/services/message_service/storage/repository"
	entity2 "pushpost/internal/services/user_service/entity"
)

type MessageUseCase struct {
	MessageRepo repository.MessageRepository
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

func (uc *MessageUseCase) GetMessagesByUserUUID(user entity2.User) (messages []entity.Message, err error) {
	messages, err = uc.MessageRepo.GetMessagesByUserUUID(user.UUID)
	if err != nil {
		return
	}
	return
}
