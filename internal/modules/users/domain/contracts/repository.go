package contracts

import "TestProject/internal/modules/users/domain/entities"

type IUsersRepository interface {
	Create(user *entities.User) error
	GetById(id string) (*entities.User, error)
	GetAll() ([]*entities.User, error)
	Update(user *entities.User, id string) error
	Delete(id string) error
}
