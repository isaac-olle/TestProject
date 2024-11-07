package create

import (
	error2 "TestProject/internal/error"
	"TestProject/internal/modules/users/domain/contracts"
	"TestProject/internal/shared/bus/domain/command"
)

type UserCreatorHandler struct {
	commandBus command.ICommandBus
	service    *UserCreator
}

func NewUserCreatorHandler(commandBus command.ICommandBus, repo contracts.IUsersRepository) {
	handler := &UserCreatorHandler{commandBus: commandBus, service: NewCreateUser(repo)}
	commandBus.RegisterHandler(CreateUserCommandName, handler.HandleCommand)
}

func (this *UserCreatorHandler) HandleCommand(command command.ICommand) error {
	cmd, ok := command.(*CreateUserCommand)
	if !ok {
		return error2.NewHttpError(500, "invalid command recieved in creator User handler")
	}
	user, err := cmd.ToDomain()
	if err != nil {
		return err
	}
	return this.service.CreateUser(user)
}
