package dto

import (
	"errors"
	"github.com/google/uuid"
	"unicode/utf8"
)

const (
	MinContentLength = 1
	MaxContentLength = 255
)

type CreateMessageDTO struct {
	SenderUUID   uuid.UUID `json:"senderUUID"`
	ReceiverUUID uuid.UUID `json:"receiverUUID"`
	Content      string    `json:"content"`
}

func (dto *CreateMessageDTO) Validate() error {
	if dto.SenderUUID == uuid.Nil {

		return errors.New("invalid sender uuid")
	}

	if dto.ReceiverUUID == uuid.Nil {

		return errors.New("invalid receiver uuid")
	}

	contentLength := utf8.RuneCountInString(dto.Content)

	if contentLength < MinContentLength || contentLength > MaxContentLength {

		return errors.New("invalid content length")
	}

	return nil
}
