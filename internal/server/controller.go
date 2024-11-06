package server

type IController interface {
	GetRoutes() []*GroupRoutes
}
