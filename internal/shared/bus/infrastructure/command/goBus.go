package command

import (
	"TestProject/internal/shared/bus/domain/command"
	errorDomain "TestProject/internal/shared/bus/domain/error"
	"fmt"
)

type GoCommandBus struct {
	asyncBus        chan command.ICommand
	handlers        map[string][]func(cmd command.ICommand) error
	errorRepository errorDomain.IErrorRepository
}

func NewGoCommandBus(errorRepository errorDomain.IErrorRepository) *GoCommandBus {
	bus := &GoCommandBus{
		asyncBus:        make(chan command.ICommand, 1024),
		handlers:        make(map[string][]func(event command.ICommand) error),
		errorRepository: errorRepository,
	}
	bus.InitializeConsumers()
	return bus
}

func (this *GoCommandBus) RegisterHandler(eventType string, f func(event command.ICommand) error) {
	this.handlers[eventType] = append(this.handlers[eventType], f)
}

func (this *GoCommandBus) Publish(event command.ICommand) {
	this.asyncBus <- event
}

func (this *GoCommandBus) Dispatch(command command.ICommand) error {
	if handlers, found := this.handlers[command.CommandType()]; found {
		for _, fnc := range handlers {
			return fnc(command)

		}
	} else {
		return fmt.Errorf("No handler found for command type: %s", command.CommandType())
	}
	return nil
}

func (this *GoCommandBus) InitializeConsumers() {
	workers := 10
	for i := 0; i < workers; i++ {
		go func(workerID int) {
			for cmd := range this.asyncBus {
				if handlers, found := this.handlers[cmd.CommandType()]; found {
					for _, fnc := range handlers {
						fmt.Printf("Worker %d processing command of type %s\n", workerID, cmd.CommandType())
						err := fnc(cmd)
						if err != nil {
							this.errorRepository.RecordError(cmd.Id(), err)
						}
					}
				} else {
					fmt.Printf("Worker %d received unknown command type: %s\n", workerID, cmd.CommandType())
				}
			}
		}(i)
	}
}
