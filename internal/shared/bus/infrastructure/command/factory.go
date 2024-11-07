package command

import (
	"TestProject/internal/config"
	"TestProject/internal/shared/bus/domain/command"
	error2 "TestProject/internal/shared/bus/domain/error"
	"fmt"
)

func GetCommandBus(config *config.DeviceConfig, repository error2.IErrorRepository) (command.ICommandBus, error) {
	switch config.Driver {
	case "go":
		return NewGoCommandBus(repository), nil
	default:
		return nil, fmt.Errorf("unsupported driver %s", config.Driver)
	}
}
