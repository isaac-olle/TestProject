package delete

import (
	"TestProject/internal/modules/users/domain/contracts"
)

type UserDeleter struct {
	repository contracts.IUsersRepository
}

func NewUserDeleter(repository contracts.IUsersRepository) *UserDeleter {
	return &UserDeleter{repository: repository}
}

func (this *UserDeleter) DeleteUser(id string) error {
	return this.repository.Delete(id)
}
