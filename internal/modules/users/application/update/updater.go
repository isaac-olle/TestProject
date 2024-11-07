package update

import (
	"TestProject/internal/modules/users/domain/contracts"
	"TestProject/internal/modules/users/domain/entities"
)

type UserUpdater struct {
	repository contracts.IUsersRepository
}

func NewUserUpdater(repository contracts.IUsersRepository) *UserUpdater {
	return &UserUpdater{repository}
}

func (this *UserUpdater) UpdateUser(user *entities.User) error {
	return this.repository.Update(user, user.Id())
}
