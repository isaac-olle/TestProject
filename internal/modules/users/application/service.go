package application

import (
	"TestProject/internal/modules/shared/domain/users"
	"TestProject/internal/modules/users/application/create"
	delete2 "TestProject/internal/modules/users/application/delete"
	"TestProject/internal/modules/users/application/get"
	"TestProject/internal/modules/users/application/patch"
	"TestProject/internal/modules/users/application/update"
	"TestProject/internal/modules/users/domain/contracts"
	"TestProject/internal/modules/users/domain/entities"
	"TestProject/internal/shared/bus/domain/command"
	"TestProject/internal/shared/bus/domain/query"
)

// El fet de que els ids ens arribin des de fora i el Create no els generi fa que l'Update i el Create necessitin exactament el mateix.
type IUserService interface {
	CreateUser(user *entities.User) error
	GetUser(userId *users.UserId) (any, error)
	UpdateUser(user *entities.User) error
	DeleteUser(id string) error
}

type UserService struct {
	*create.UserCreator
	*get.UserGetter
	*update.UserUpdater
	*patch.UserPatcher
	*delete2.UserDeleter
}

func NewUserService(commandBus command.ICommandBus, queryBus query.IQueryBus, repository contracts.IUsersRepository) {
	create.NewUserCreatorHandler(commandBus, repository)
	get.NewUserGetterHandler(queryBus, repository)
	update.NewUserUpdaterHandler(commandBus, repository)
}
