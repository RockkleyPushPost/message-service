package repository

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"pushpost/internal/services/message_service/entity"
)

type MessageRepository struct {
	DB *gorm.DB
}

func NewMessageRepository(DB *gorm.DB) *MessageRepository {
	return &MessageRepository{DB: DB}
}

func (r *MessageRepository) CreateMessage(message *entity.Message) error { // FIXME WTF
	if r.DB.Find(&entity.Message{}, "uuid = ?", message.ReceiverUUID).Error != nil {

		return errors.New("receiver not found")
	}

	if r.DB.Find(&entity.Message{}, "uuid = ?", message.SenderUUID).Error != nil {

		return errors.New("sender not found")
	}

	return r.DB.Create(&message).Error
}

func (r *MessageRepository) GetMessagesByUserUUID(uuid uuid.UUID) (messages []entity.Message, err error) {
	err = r.DB.Where("sender_uuid = ? OR receiver_uuid = ?", uuid, uuid).Find(&messages).Error

	return
}
