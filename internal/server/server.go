package server

import (
	"TestProject/internal/config"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Server struct {
	controllers []IController
	*gin.Engine
}

func NewServer() *Server {
	return &Server{nil, gin.Default()}
}

func (this *Server) AddController(controller IController) {
	this.controllers = append(this.controllers, controller)
}

func (this *Server) Run(config *config.BasicConfig) error {
	for _, controller := range this.controllers {
		this.RegisterEndpoints(controller)
	}
	return this.Engine.Run(fmt.Sprintf(":%s", config.Port))
}

func (this *Server) RegisterEndpoints(controller IController) {
	for _, v := range controller.GetRoutes() {
		group := this.Group(v.Group)
		for _, m := range v.Middlewares {
			group.Use(m)
		}
		for _, v2 := range v.Routes {
			group.Handle(v2.Method, v2.Path, v2.HandlerFunc)
		}
	}
}
