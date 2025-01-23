package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	UUID         uuid.UUID
	SenderUUID   uuid.UUID
	ReceiverUUID uuid.UUID
	Content      string
}
