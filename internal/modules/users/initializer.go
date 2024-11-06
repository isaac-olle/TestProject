package users

import (
	"TestProject/internal/config"
	"TestProject/internal/modules/users/application"
	"TestProject/internal/modules/users/infrastructure/http"
	"TestProject/internal/modules/users/infrastructure/persistance"
	"TestProject/internal/server"
	"TestProject/internal/shared/bus/domain/command"
	"TestProject/internal/shared/bus/domain/query"
)

type UserInitializer struct {
	controller server.IController
}

func NewUserInitializer(persistanceConfig *config.DeviceConfig, commandBus command.ICommandBus, queryBus query.IQueryBus) (*UserInitializer, error) {
	repo, err := persistance.NewUsersRepository(persistanceConfig)
	if err != nil {
		return nil, err
	}
	application.NewUserHandler(commandBus, queryBus, repo)
	controller := http.NewUsersController(commandBus, queryBus)
	return &UserInitializer{
		controller: controller,
	}, nil
}

func (this *UserInitializer) Controller() server.IController {
	return this.controller
}
