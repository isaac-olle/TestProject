package application

import (
	error2 "TestProject/internal/error"
	"TestProject/internal/modules/users/application/create"
	"TestProject/internal/modules/users/application/get"
	"TestProject/internal/modules/users/domain/contracts"
	"TestProject/internal/shared/bus/domain/command"
	"TestProject/internal/shared/bus/domain/query"
)

type UserHandler struct {
	commandBus command.ICommandBus
	queryBus   query.IQueryBus
	service    IUserService
}

func NewUserHandler(commandBus command.ICommandBus, queryBus query.IQueryBus, repo contracts.IUsersRepository) {
	service := NewUserService(repo)
	handler := &UserHandler{commandBus: commandBus, queryBus: queryBus, service: service}
	commandBus.RegisterHandler(create.UserCreatedCommand, handler.HandleCommands)
	queryBus.RegisterHandler(get.GetUserQueryName, handler.HandleQueries)
}

// Els casos d'ús han de rebre l'objecte ja creat, per aixo aqui es transforma als VO. Si no tinguessim CQRS l'encarregat de fer aixo seria el Controller. Explicat al curs.
func (this *UserHandler) HandleCommands(command command.ICommand) error {
	switch cmd := command.(type) {
	case *create.CreateUserCommand:
		user, err := cmd.ToDomain()
		if err != nil {
			return err
		}
		return this.service.CreateUser(user)
	default:
		return error2.NewHttpError(500, "invalid command recieved in User handler")
	}
}

// Aquest any es un JSON
func (this *UserHandler) HandleQueries(query query.IQuery) (any, error) {
	switch cmd := query.(type) {
	case *get.GetUserQuery:
		return this.service.GetUser(cmd.Id())
	default:
		return nil, error2.NewHttpError(500, "invalid query recieved in User handler")
	}
}
