package server

import "TestProject/internal/config"

type IServer interface {
	Run(config *config.BasicConfig) error
	AddController(controller IController)
}
