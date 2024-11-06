package create

import (
	"TestProject/internal/modules/users/domain/contracts"
	"TestProject/internal/modules/users/domain/entities"
)

type UserCreator struct {
	repository contracts.IUsersRepository
}

func NewCreateUser(repository contracts.IUsersRepository) *UserCreator {
	return &UserCreator{repository: repository}
}

func (this *UserCreator) CreateUser(user *entities.User) error {
	err := this.repository.Create(user)
	if err != nil {
		return err
	}
	return nil
}
