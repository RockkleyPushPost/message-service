package di

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"pushpost/internal/services/message_service/transport"
)

type Container struct {
	Server  *fiber.App
	DB      *gorm.DB
	Handler transport.MessageHandler
}

func NewContainer() (*Container, error) {
	container := &Container{}

	//messageRepo.DB.AutoMigrate(entity.Message{}) //fixme make goose migrations
	//userRepo.DB.AutoMigrate(entity.User{})       //fixme make goose migrations

	return container, nil
}

func (c *Container) RegisterHandler(handlerGroup transport.MessageHandler) {
	c.Handler = handlerGroup
}
