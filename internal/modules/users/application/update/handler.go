package update

import (
	error2 "TestProject/internal/error"
	"TestProject/internal/modules/users/domain/contracts"
	"TestProject/internal/shared/bus/domain/command"
)

type UserUpdaterHandler struct {
	commandBus command.ICommandBus
	service    *UserUpdater
}

func NewUserUpdaterHandler(commandBus command.ICommandBus, repo contracts.IUsersRepository) {
	handler := &UserUpdaterHandler{commandBus: commandBus, service: NewUserUpdater(repo)}
	commandBus.RegisterHandler(UpdateUserCommandName, handler.HandleCommand)
}

func (this *UserUpdaterHandler) HandleCommand(command command.ICommand) error {
	cmd, ok := command.(*UpdateUserCommand)
	if !ok {
		return error2.NewHttpError(500, "invalid command recieved in updater User handler")
	}
	user, err := cmd.ToDomain()
	if err != nil {
		return err
	}
	return this.service.UpdateUser(user)
}
