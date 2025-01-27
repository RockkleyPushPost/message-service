package di

import (
	"pushpost/internal/config"
	"pushpost/internal/services/message_service/domain"
	"pushpost/internal/services/message_service/domain/usecase"
	"pushpost/internal/services/message_service/storage"
	"pushpost/internal/services/message_service/storage/repository"
	"pushpost/internal/services/message_service/transport"
	transport2 "pushpost/internal/services/message_service/transport/handlers"
	"pushpost/internal/services/message_service/transport/routing"
	"pushpost/internal/setup"
)

func Setup(cfg config.Config, di *Container) error {
	db, err := setup.Database(cfg.Database)
	di.DB = db
	if err != nil {
		return err
	}

	fiber := setup.NewFiber()
	di.Server = fiber
	var messageRepository storage.MessageRepository = repository.NewMessageRepository(db)
	var messageUseCase domain.MessageUseCase = usecase.NewMessageUseCase(messageRepository)
	var messageHandler transport.MessageHandler = transport2.NewMessagesHandler(messageUseCase, fiber)

	routing.SetupRoutes(messageHandler)
	di.RegisterHandler(messageHandler)

	return nil
}

//
//func Setup(cfg config.Config, di *di.Container) error {
//	db, err := setup.Database(&cfg.Database)
//
//	if err != nil {
//		return err
//	}
//
//	di.Register(db)
//
//	fiber := setup.NewFiber(&cfg.Fiber)
//
//	di.Register(fiber)
//
//	// Domain
//
//	// Domain - UseCases
//
//	var messageUseCase domain.MessageUseCase = &usecase.MessageUseCase{}
//	di.Register(messageUseCase)
//
//	// Storage
//
//	// Storage - Repositories
//
//	var messageRepository storage.MessageRepository = &repository.MessageRepository{}
//	di.Register(messageRepository)
//
//	// Routing
//
//	var messageHandler transport.MessageHandler = &transport2.MessagesHandler{}
//
//	routing.SetupRoutes(messageHandler)
//	di.Register(messageHandler)
//
//	return nil
//}
