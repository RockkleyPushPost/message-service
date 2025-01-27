package service

import (
	"context"
	"errors"
	"log"
	"pushpost/internal/config"
	"pushpost/internal/services/message_service/service/di"
)

const servName string = "MessageService"

type Option func(*service)

type service struct {
	container *di.Container
	logger    *log.Logger
	config    *config.Config
}

func NewService(opts ...Option) (Service, error) {
	s := &service{}
	for _, opt := range opts {
		opt(s)
	}

	return s, s.validate()
}

func WithContainer(container *di.Container) Option {
	return func(s *service) {
		s.container = container
	}
}

func WithLogger(logger *log.Logger) Option {
	return func(s *service) {
		s.logger = logger
	}
}

func WithConfig(config *config.Config) Option {
	return func(s *service) {
		s.config = config
	}
}

func (s *service) validate() error {

	if s.container == nil {
		return errors.New("missing container")
	}
	if s.logger == nil {
		return errors.New("missing logger")
	}
	if s.config == nil {
		return errors.New("missing config")
	}
	return nil
}
func (s *service) Run(ctx context.Context) error {
	s.logger.Printf("starting %s\n", s.Name())

	err := s.container.Server.Listen(":3001") //fixme gotta parse
	if err != nil {
		return err
	}
	<-ctx.Done()
	return nil
}

func (s *service) Shutdown(ctx context.Context) error {
	s.logger.Printf("shutting down %s...\n", servName)
	// todo
	return nil
}

func (s *service) Name() string {
	return servName
}
