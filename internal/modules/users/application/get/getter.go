package get

import (
	"TestProject/internal/modules/users/domain/contracts"
)

type UserGetter struct {
	repository contracts.IUsersRepository
}

func NewUserGetter(repository contracts.IUsersRepository) *UserGetter {
	return &UserGetter{repository: repository}
}

func (this *UserGetter) GetUser(id string) (any, error) {
	user, err := this.repository.GetById(id)
	if err != nil {
		return nil, err
	}
	return user.ToResponse(), user.CheckDeleted()
}
