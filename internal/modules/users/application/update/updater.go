package update

import "TestProject/internal/modules/users/domain/contracts"

type UserUpdater struct {
	repository contracts.IUsersRepository
}

func NewUserUpdater(repository contracts.IUsersRepository) *UserUpdater {
	return &UserUpdater{repository}
}

func (this *UserUpdater) UpdateUser() {
}
