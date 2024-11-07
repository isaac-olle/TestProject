package command

import (
	"TestProject/internal/config"
	error2 "TestProject/internal/shared/bus/domain/error"
	"TestProject/internal/shared/bus/domain/query"
	"fmt"
)

func GetQueryBus(config *config.DeviceConfig, repository error2.IErrorRepository) (query.IQueryBus, error) {
	switch config.Driver {
	case "go":
		return NewGoQueryBus(repository), nil
	default:
		return nil, fmt.Errorf("unsupported driver %s", config.Driver)
	}
}
