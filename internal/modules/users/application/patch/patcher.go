package patch

import "TestProject/internal/modules/users/domain/contracts"

type UserPatcher struct {
	repository contracts.IUsersRepository
}

func NewUserPatcher(repository contracts.IUsersRepository) *UserPatcher {
	return &UserPatcher{repository: repository}
}

func (this *UserPatcher) PatchUser() {}
