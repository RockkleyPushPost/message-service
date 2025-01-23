package storage

import (
	"github.com/google/uuid"
	"pushpost/internal/services/message_service/entity"
)

type MessageRepository interface {
	CreateMessage(message *entity.Message) error
	GetMessagesByUserUUID(uuid uuid.UUID) (messages []entity.Message, err error)
}
