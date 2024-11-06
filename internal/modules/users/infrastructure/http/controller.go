package http

import (
	"TestProject/internal/modules/users/application/create"
	"TestProject/internal/modules/users/application/get"
	"TestProject/internal/server"
	"TestProject/internal/shared/bus/domain/command"
	"TestProject/internal/shared/bus/domain/query"
	"TestProject/internal/shared/uuid/domain"
	"TestProject/internal/shared/uuid/infrastructure"
	"github.com/gin-gonic/gin"
	"net/http"
)

const identifier = "id"

type UsersController struct {
	commandBus  command.ICommandBus
	queryBus    query.IQueryBus
	idGenerator domain.IIdGenerator
}

func NewUsersController(commandBus command.ICommandBus, queryBus query.IQueryBus) *UsersController {
	//Aixo hauria d'entrar per dependencies, pero per no liarho mes ho fem aqui
	return &UsersController{
		commandBus:  commandBus,
		queryBus:    queryBus,
		idGenerator: infrastructure.NewUuidGenerator(),
	}
}

func (this *UsersController) GetRoutes() []*server.GroupRoutes {
	return []*server.GroupRoutes{
		{
			Group:       "/users",
			Middlewares: []gin.HandlerFunc{server.ContextMiddleware, server.ErrorHandlerMiddleware},
			Routes: []*gin.RouteInfo{
				{
					Method:      "POST",
					Path:        "",
					Handler:     "users.CreateUser",
					HandlerFunc: this.CreateUser,
				},
				{
					Method:      "GET",
					Path:        "/:id",
					Handler:     "users.GetUser",
					HandlerFunc: this.GetUser,
				}, {
					Method:      "PUT",
					Path:        "/:id",
					Handler:     "users.UpdateUser",
					HandlerFunc: this.UpdateUser,
				},
				{
					Method:      "DELETE",
					Path:        "/:id",
					Handler:     "users.DeleteUser",
					HandlerFunc: this.DeleteUser,
				},
			},
		},
		{
			Group:       "/users/async",
			Middlewares: []gin.HandlerFunc{server.ContextMiddleware, server.ErrorHandlerMiddleware},
			Routes: []*gin.RouteInfo{
				{
					Method:      "POST",
					Path:        "",
					Handler:     "users.CreateUser",
					HandlerFunc: this.CreateUser,
				},
				{
					Method:      "GET",
					Path:        "/:id",
					Handler:     "users.GetUser",
					HandlerFunc: this.GetUser,
				}, {
					Method:      "PUT",
					Path:        "/:id",
					Handler:     "users.UpdateUser",
					HandlerFunc: this.UpdateUser,
				},
				{
					Method:      "DELETE",
					Path:        "/:id",
					Handler:     "users.DeleteUser",
					HandlerFunc: this.DeleteUser,
				},
			},
		},
	}
}

func (this *UsersController) CreateUser(c *gin.Context) {
	var auxId string
	var createUserCommand *create.CreateUserCommand
	if err := c.ShouldBindJSON(&createUserCommand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "request body does not fit with the expected. Gin Error:" + err.Error()})
		return
	}
	if createUserCommand.AbstractCommand == nil {
		auxId = this.idGenerator.Generate().String()
		createUserCommand.AbstractCommand = command.NewAbstractCommand(auxId)
	}
	err := this.commandBus.Dispatch(createUserCommand)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": auxId})
	return
}

func (this *UsersController) GetUser(c *gin.Context) {
	id := c.Param(identifier)
	resp, err := this.queryBus.Dispatch(get.NewGetUserQuery(id))
	if err != nil {
		c.Error(err)
	} else {
		c.JSON(http.StatusCreated, resp)
	}
	return
}

func (this *UsersController) UpdateUser(c *gin.Context) {
	return
}

func (this *UsersController) PatchUser(c *gin.Context) {
	return
}

func (this *UsersController) DeleteUser(c *gin.Context) {
	return
}

func (this *UsersController) CreateUserAsync(c *gin.Context) {
	var auxId string
	var createUserCommand *create.CreateUserCommand
	if err := c.ShouldBindJSON(&createUserCommand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "request body does not fit with the expected. Gin Error:" + err.Error()})
		return
	}
	if createUserCommand.AbstractCommand == nil {
		auxId = this.idGenerator.Generate().String()
		createUserCommand.AbstractCommand = command.NewAbstractCommand(auxId)
	}
	this.commandBus.Publish(createUserCommand)
	c.JSON(http.StatusCreated, gin.H{"id": auxId})
}

func (this *UsersController) GetUserAsync(c *gin.Context) {
	return
}

func (this *UsersController) UpdateUserAsync(c *gin.Context) {
	return
}

func (this *UsersController) PatchUserAsync(c *gin.Context) {
	return
}

func (this *UsersController) DeleteUserAsync(c *gin.Context) {
	return
}
