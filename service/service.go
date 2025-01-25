package service

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const servName string = "MessageService"

type Service struct {
	Database *gorm.DB
	Server   *fiber.App
}

func (*Service) Name() string {
	return servName
}
