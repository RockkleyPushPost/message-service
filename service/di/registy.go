package di

type HandlerConstructor func(container *Container) interface{}

type HandlerRegistry struct {
	handlers map[string]HandlerConstructor
}

func NewHandlerRegistry() *HandlerRegistry {
	return &HandlerRegistry{handlers: make(map[string]HandlerConstructor)}
}

func (r *HandlerRegistry) Registry(name string, constructor HandlerConstructor) {
	r.handlers[name] = constructor
}
