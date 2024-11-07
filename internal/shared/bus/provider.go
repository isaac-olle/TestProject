package bus

import (
	"TestProject/internal/config"
	"TestProject/internal/shared/bus/domain/command"
	"TestProject/internal/shared/bus/domain/query"
	command2 "TestProject/internal/shared/bus/infrastructure/command"
	query2 "TestProject/internal/shared/bus/infrastructure/query"
)

type BusProvider struct {
	commandBus command.ICommandBus
	queryBus   query.IQueryBus
}

func NewBusProvider(errorRepoConfig *config.DeviceConfig, commandBusConfig *config.DeviceConfig, queryBusConfig *config.DeviceConfig) (*BusProvider, error) {
	/*errorRepository, err := error2.NewErrorRepository(errorRepoConfig)
	if err != nil {
		return nil, err
	}*/
	commandBus, err := command2.GetCommandBus(commandBusConfig, nil)
	if err != nil {
		return nil, err
	}
	queryBus, err := query2.GetQueryBus(queryBusConfig, nil)
	if err != nil {
		return nil, err
	}
	return &BusProvider{commandBus, queryBus}, nil
}

func (this *BusProvider) CommandBus() command.ICommandBus {
	return this.commandBus
}

func (this *BusProvider) QueryBus() query.IQueryBus {
	return this.queryBus
}
