package contracts

import (
	"TestProject/internal/modules/shared/domain/users"
	"TestProject/internal/modules/users/domain/entities"
)

type IUsersRepository interface {
	Create(user *entities.User) error
	GetById(id *users.UserId) (*entities.User, error)
	GetAll() ([]*entities.User, error)
	Update(user *entities.User, id *users.UserId) error
	Delete(id string) error
}
