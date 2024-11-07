package get

import (
	error2 "TestProject/internal/error"
	"TestProject/internal/modules/shared/domain/users"
	"TestProject/internal/modules/users/domain/contracts"
	"TestProject/internal/shared/bus/domain/query"
)

type UserGetterHandler struct {
	queryBus query.IQueryBus
	service  *UserGetter
}

func NewUserGetterHandler(queryBus query.IQueryBus, repo contracts.IUsersRepository) {
	handler := &UserGetterHandler{queryBus: queryBus, service: NewUserGetter(repo)}
	queryBus.RegisterHandler(GetUserQueryName, handler.HandleCommand)
}

func (this *UserGetterHandler) HandleCommand(query query.IQuery) (any, error) {
	cmd, ok := query.(*GetUserQuery)
	if !ok {
		return nil, error2.NewHttpError(500, "invalid command recieved in updater User handler")
	}
	userId, err := users.NewUserIdFromString(cmd.Id())
	if err != nil {
		return nil, error2.NewBadRequestHttpError("invalid user id")
	}
	return this.service.GetUser(userId)
}
