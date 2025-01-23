package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"pushpost/internal/services/message_service/domain/dto"
)

type Message struct {
	gorm.Model
	UUID         uuid.UUID
	SenderUUID   uuid.UUID
	ReceiverUUID uuid.UUID
	Content      string
}

func NewMessage(dto dto.CreateMessageDTO) *Message {
	return &Message{
		UUID:         uuid.New(),
		SenderUUID:   dto.SenderUUID,
		ReceiverUUID: dto.ReceiverUUID,
		Content:      dto.Content,
	}
}
