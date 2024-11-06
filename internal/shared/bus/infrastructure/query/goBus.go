package command

import (
	error2 "TestProject/internal/error"
	errorDomain "TestProject/internal/shared/bus/domain/error"
	"TestProject/internal/shared/bus/domain/query"
	"errors"
	"fmt"
)

type GoQueryBus struct {
	handlers        map[string][]func(cmd query.IQuery) (any, error)
	errorRepository errorDomain.IErrorRepository
}

func NewGoQueryBus(errorRepository errorDomain.IErrorRepository) *GoQueryBus {
	return &GoQueryBus{
		handlers:        make(map[string][]func(event query.IQuery) (any, error)),
		errorRepository: errorRepository,
	}
}

func (this *GoQueryBus) RegisterHandler(eventType string, f func(event query.IQuery) (any, error)) {
	this.handlers[eventType] = append(this.handlers[eventType], f)
}

func (this *GoQueryBus) Dispatch(query query.IQuery) (any, error) {
	if handlers, found := this.handlers[query.QueryType()]; found {
		for _, fnc := range handlers {
			var notFoundError *error2.NotFoundHttpError
			resp, err := fnc(query)
			if err == nil {
				return resp, nil
			}
			if errors.Is(err, notFoundError) {
				return this.errorRepository.GetError(query.Id())
			}
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("no handler found for command type: %s", query)
	}
	return nil, nil
}
