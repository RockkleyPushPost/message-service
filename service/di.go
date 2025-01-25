package service

import (
	"pushpost/internal/di"
	"pushpost/internal/services/message_service/config"
	"pushpost/internal/services/message_service/domain"
	"pushpost/internal/services/message_service/domain/usecase"
	"pushpost/internal/services/message_service/storage"
	"pushpost/internal/services/message_service/storage/repository"
	"pushpost/internal/setup"
)

func Setup(cfg config.Config, di *di.Container) error {
	db, err := setup.Database(&cfg.Database)

	if err != nil {
		return err
	}

	di.Register(db)

	fiber := setup.NewFiber(&cfg.Fiber)

	di.Register(fiber)

	// Domain

	// Domain - UseCases

	var messageUseCase domain.MessageUseCase = &usecase.MessageUseCase{}
	di.Register(messageUseCase)

	// Storage

	// Storage - Repositories

	var messageRepository storage.MessageRepository = &repository.MessageRepository{}
	di.Register(messageRepository)

	return nil
}
