package application

import (
	"TestProject/internal/modules/users/application/create"
	delete2 "TestProject/internal/modules/users/application/delete"
	"TestProject/internal/modules/users/application/get"
	"TestProject/internal/modules/users/application/patch"
	"TestProject/internal/modules/users/application/update"
	"TestProject/internal/modules/users/domain/contracts"
	"TestProject/internal/modules/users/domain/entities"
)

type IUserService interface {
	CreateUser(user *entities.User) error
	GetUser(userId string) (any, error)
	DeleteUser(id string) error
}

type UserService struct {
	*create.UserCreator
	*get.UserGetter
	*update.UserUpdater
	*patch.UserPatcher
	*delete2.UserDeleter
}

func NewUserService(repository contracts.IUsersRepository) IUserService {
	return &UserService{
		create.NewCreateUser(repository),
		get.NewUserGetter(repository),
		update.NewUserUpdater(repository),
		patch.NewUserPatcher(repository),
		delete2.NewUserDeleter(repository),
	}
}
