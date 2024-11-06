package server

import "github.com/gin-gonic/gin"

type GroupRoutes struct {
	Group       string
	Middlewares []gin.HandlerFunc
	Routes      []*gin.RouteInfo
}
