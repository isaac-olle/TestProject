package domain

// El Sync bus, depen de si es command o query necessita retornar o error o []byte, i error de manera que no pot ser sense generics
type IAsyncBus[T IEvent] interface {
	Publish(event T)
	InitializeConsumers()
}
