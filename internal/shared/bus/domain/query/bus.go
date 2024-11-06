package query

type IQueryBus interface {
	RegisterHandler(eventType string, f func(query IQuery) (any, error))
	Dispatch(event IQuery) (any, error)
}
