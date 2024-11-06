package bus

import (
	"TestProject/internal/config"
	"TestProject/internal/shared/bus/domain/command"
	"TestProject/internal/shared/bus/domain/query"
	command2 "TestProject/internal/shared/bus/infrastructure/command"
	error2 "TestProject/internal/shared/bus/infrastructure/error"
	query2 "TestProject/internal/shared/bus/infrastructure/query"
)

type BusProvider struct {
	commandBus command.ICommandBus
	queryBus   query.IQueryBus
}

func NewBusProvider(errorRepoConfig *config.DeviceConfig, commandBusConfig *config.DeviceConfig, queryBusConfig *config.DeviceConfig) (*BusProvider, error) {
	errorRepository, err := error2.NewErrorRepository(errorRepoConfig)
	if err != nil {
		return nil, err
	}
	return &BusProvider{command2.NewGoCommandBus(errorRepository), query2.NewGoQueryBus(errorRepository)}, nil
}

func (this *BusProvider) CommandBus() command.ICommandBus {
	return this.commandBus
}

func (this *BusProvider) QueryBus() query.IQueryBus {
	return this.queryBus
}
