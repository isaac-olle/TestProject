package modules

import (
	"TestProject/internal/config"
	"TestProject/internal/modules/users"
	"TestProject/internal/server"
	"TestProject/internal/shared/bus/domain/command"
	"TestProject/internal/shared/bus/domain/query"
)

type ModuleInitializer struct {
	users *users.UserInitializer
}

func NewModuleInitializer(persistanceConfig *config.DeviceConfig, commandBus command.ICommandBus, queryBus query.IQueryBus) (*ModuleInitializer, error) {
	userInitializer, err := users.NewUserInitializer(persistanceConfig, commandBus, queryBus)
	if err != nil {
		return nil, err
	}
	return &ModuleInitializer{users: userInitializer}, nil
}

func (this *ModuleInitializer) AddControllers(server server.IServer) {
	server.AddController(this.users.Controller())
}
