package create

import (
	"TestProject/internal/modules/users/domain/dtos"
	"TestProject/internal/modules/users/domain/entities"
	"TestProject/internal/shared/bus/domain/command"
)

const (
	birthdateTimeFormat = "2006-01-02"
	UserCreatedCommand  = "user_created"
)

type CreateUserCommand struct {
	Name      string `json:"name" binding:"required"`
	Surname   string `json:"surname,omitempty"`
	BirthDate string `json:"birthDate"`
	Email     string `json:"email" binding:"required,email"`

	*command.AbstractCommand
}

func (req *CreateUserCommand) ToDomain() (*entities.User, error) {
	return entities.NewUserFromUnvaluedObjects(dtos.UserParams{
		Id:                  req.ID,
		Name:                req.Name,
		Surname:             req.Surname,
		Email:               req.Email,
		Birthdate:           req.BirthDate,
		BirthdateTimeFormat: birthdateTimeFormat,
		CreatedAt:           "",
		CreatedAtTimeFormat: "",
	})
}

func (this *CreateUserCommand) CommandType() string {
	return UserCreatedCommand
}
