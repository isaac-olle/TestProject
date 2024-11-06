package command

import "TestProject/internal/shared/bus/domain"

type ICommandBus interface {
	domain.IAsyncBus[ICommand]
	RegisterHandler(eventType string, f func(iCommand ICommand) error)
	Dispatch(event ICommand) error
}
